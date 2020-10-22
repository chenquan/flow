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
	"os"
	"sync"
)

var _ Data = (*os.File)(nil)

// ChanData 数据流通道
type ChanData chan Data

// ResultFunc 流结果处理函数
type ResultFunc func(result Data)

// Flow 流
type Flow struct {
	root Node
	in   chan Data
	out  chan Data
	buff int
	wg   sync.WaitGroup
}

// NewFlow 新建一条流处理
func NewFlow(buff int) *Flow {
	f := func(in Data) Data {
		return in
	}
	return &Flow{
		root: &FuncNode{
			funcNode: f,
		},
		in:   make(chan Data, buff),
		out:  make(chan Data, buff),
		buff: buff}
}

// FlowIn 数据流入流节点
func (f *Flow) FlowIn(node Node) Node {
	f.root.FlowIn(node)
	return node
}

// FlowInWithFunc 数据流入函数流节点
func (f *Flow) FlowInWithFunc(funcNode Func) Node {
	node := NewFuncNode(funcNode)
	f.root.FlowIn(node)
	return node
}

// Run 建立流处理通道
func (f *Flow) Run() {
	node := f.root

	nodeChans := make([]ChanData, 0)
	nodeChans = append(nodeChans, f.in)
	for node != nil && node.Next() != nil {
		nodeChans = append(nodeChans, make(ChanData, f.buff))
		node = node.Next()
	}
	nodeChans = append(nodeChans, f.out)

	node = f.root
	for i := 0; node != nil; i++ {
		in := nodeChans[i]
		out := nodeChans[i+1]
		go func(node Node) {
			for data := range in {
				go func(data Data) {
					out <- node.Run(data)
				}(data)
			}
		}(node)
		node = node.Next()
	}

}

// Feed 喂入流处理数据
func (f *Flow) Feed(inData Data, resultFunc ResultFunc) {
	f.wg.Add(1)
	f.in <- inData
	go func(resultFunc func(inData Data)) {
		resultFunc(<-f.out)
		f.wg.Done()
	}(resultFunc)

}

// Wait 等待全部流结束
func (f *Flow) Wait() {
	f.wg.Wait()
}