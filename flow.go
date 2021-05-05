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

// ChanContext 数据流通道
type ChanContext chan *Context

// ResultFunc 流结果处理函数
//type ResultFunc func(ctx *Context)

// Flow 流
type Flow struct {
	root Node
	in   chan *Context
	out  chan *Context
	buff int
	wg   sync.WaitGroup
	pool *ants.PoolWithFunc
}

// NewFlow 新建一条流处理
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

// ToNode 数据流入流节点
func (f *Flow) ToNode(node Node) Node {
	f.root.ToNode(node)
	return node
}

// To 数据流入函数流节点
func (f *Flow) To(funcNode Func) Node {
	node := f.root.To(funcNode)
	return node
}

type NodeData struct {
	node Node
	ctx  *Context
	out  ChanContext
}

// Run 建立流处理通道
func (f *Flow) Run(pool bool) {
	fn := func(ctx *Context, node Node, out ChanContext) {
		// 确保每个协程执行完毕
		f.wg.Add(1)
		defer f.wg.Done()
		node.Run(ctx)
		if ctx.Err() == nil {
			out <- ctx
		} else {
			// 将错误信息发送给输出通道
			f.out <- ctx
		}
		atomic.AddInt32(&ctx.step, 1)
	}
	// 协程池
	f.pool, _ = ants.NewPoolWithFunc(1000, func(c interface{}) {
		nodeData := c.(*NodeData)
		fn(nodeData.ctx, nodeData.node, nodeData.out)
	})

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
		go func(node Node) {
			for ctx := range in {
				if pool {
					err := f.pool.Invoke(&NodeData{
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
		}(node)
		node = node.Next()
	}

}

// Feed 喂入流处理数据
func (f *Flow) Feed(data interface{}, resultFunc Func) string {
	f.wg.Add(1)
	ctx := NewContext(data)
	f.in <- ctx
	go func(resultFunc func(inData *Context)) {
		resultFunc(<-f.out)
		f.wg.Done()
	}(resultFunc)
	return ctx.FlowId()
}

// FeedData 喂入流处理数据
func (f *Flow) FeedData(ctx *Context, resultFunc Func) string {
	f.wg.Add(1)
	f.in <- ctx
	go func(resultFunc func(ctx *Context)) {
		resultFunc(<-f.out)
		f.wg.Done()
	}(resultFunc)
	return ctx.FlowId()
}

// Wait 等待全部流结束
func (f *Flow) Wait() {
	f.wg.Wait()
	defer f.pool.Release()
}
