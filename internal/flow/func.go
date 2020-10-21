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
	"context"
	"sync"
)

type Func func(in Data) Data

var _ Node = (*FuncNode)(nil)

// FuncNode 函数计算流节点
type FuncNode struct {
	funcNode      Func
	flow          bool // 是否已运行
	nextFuncFlows Node // 子计算节点

	ctx        context.Context
	cancelFunc context.CancelFunc
	mu         sync.RWMutex
}

func (f *FuncNode) FlowInWithFunc(funcNode Func) Node {
	node := NewFuncNode(funcNode)
	f.nextFuncFlows = node
	return node
}

func NewFuncNode(funcNode Func) Node {
	return &FuncNode{funcNode: funcNode}
}

func (f *FuncNode) FlowIn(node Node) Node {
	//ctx, cancelFunc := context.WithCancel(f.ctx)
	//node.SetParentContext(ctx, cancelFunc)
	f.nextFuncFlows = node
	return node
}

func (f *FuncNode) Next() Node {
	return f.nextFuncFlows
}

func (f *FuncNode) Run(in Data) (out Data) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.flow = true
	out = f.funcNode(in)
	return
}

func (f *FuncNode) SetParentContext(ctx context.Context, cancelFunc context.CancelFunc) {
	f.ctx = ctx
	f.cancelFunc = cancelFunc
}
