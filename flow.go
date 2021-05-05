/*
 *
 *     Copyright 2020 yunqi
 *
 *     Licensed under the Apache License, Version 2.0 (the "License");
 *     you may not use this file except in compliance with the License.
 *     You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 *     Unless required by applicable law or agreed to in writing, software
 *     distributed under the License is distributed on an "AS IS" BASIS,
 *     WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *     See the License for the specific language governing permissions and
 *     limitations under the License.
 *
 */

package flow

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"sync/atomic"
)

// ChanContext Data stream channel
type ChanContext chan *Context

// Flow Implements a flow
type Flow struct {
	root  Node
	in    chan *Context
	out   chan *Context
	buff  int
	wg    sync.WaitGroup
	pools []*ants.PoolWithFunc
}

// NewFlow Create a new stream processing
func NewFlow(buff int) *Flow {
	f := func(in *Context) {

	}
	return &Flow{
		root: &FuncNode{
			funcNode: f,
		},
		in:   make(chan *Context, buff),
		out:  make(chan *Context, buff),
		buff: buff}
}

// ToNode Data flows into the flow node
func (f *Flow) ToNode(node Node) Node {
	f.root.ToNode(node)
	return node
}

// To Data flows into the function flow node
func (f *Flow) To(funcNode Func) Node {
	node := f.root.To(funcNode)
	return node
}

// NodeData  Node data
type NodeData struct {
	node Node
	ctx  *Context
	out  ChanContext
}

// Run Establish a stream processing channel
func (f *Flow) Run(resultFunc Func, options ...Option) {

	// option
	op := loadOptions(options...)

	fn := func(ctx *Context, node Node, out ChanContext) {
		// Ensure that each coroutine is executed
		f.wg.Add(1)
		defer f.wg.Done()
		node.Run(ctx)
		if ctx.Err() == nil {
			out <- ctx
		} else {
			// Send error information to the output channel
			f.out <- ctx
		}
		atomic.AddInt32(&ctx.step, 1)
	}

	node := f.root

	nodeChans := make([]ChanContext, 0)
	nodeChans = append(nodeChans, f.in)
	for node != nil && node.Next() != nil {
		nodeChans = append(nodeChans, make(ChanContext, f.buff))
		node = node.Next()
	}
	nodeChans = append(nodeChans, f.out)

	node = f.root
	for i := 0; node != nil; i++ {
		in := nodeChans[i]
		out := nodeChans[i+1]
		// Node goroutine pool
		var nodePool *ants.PoolWithFunc
		if !op.disablePool {
			nodePool, _ = ants.NewPoolWithFunc(op.poolSize, func(c interface{}) {
				nodeData := c.(*NodeData)
				fn(nodeData.ctx, nodeData.node, nodeData.out)
			})
			f.pools = append(f.pools, nodePool)
		}
		// execute node
		go func(node Node, nodePool *ants.PoolWithFunc, in, out ChanContext, i int) {
			for ctx := range in {
				if !op.disablePool {
					err := nodePool.Invoke(&NodeData{
						node: node,
						ctx:  ctx,
						out:  out,
					})
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fn(ctx, node, out)
				}
			}
		}(node, nodePool, in, out, i)

		node = node.Next()
	}
	// process flow results
	go func() {
		// result goroutine pool
		resultPool, _ := ants.NewPoolWithFunc(op.poolSize, func(c interface{}) {
			ctx := c.(*Context)
			resultFunc(ctx)
			f.wg.Done()
		})
		f.pools = append(f.pools, resultPool)
		for ctx := range f.out {
			err := resultPool.Invoke(ctx)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()

}

// Feed Feed stream processing data
func (f *Flow) Feed(data interface{}) string {
	f.wg.Add(1)
	ctx := newContext(data)
	f.in <- ctx
	return ctx.FlowId()
}

// Wait Wait for all streams to end
func (f *Flow) Wait() {
	f.wg.Wait()
	defer func() {
		// 释放协程池
		if f.pools != nil {
			for _, pool := range f.pools {
				pool.Release()
			}
		}
	}()
}
