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

// Func 节点处理函数
type Func func(ctx *Context)

var _ Node = (*FuncNode)(nil)

// FuncNode 函数流节点
type FuncNode struct {
	funcNode      Func
	nextFuncFlows Node // 子计算节点

}

// To 数据流入函数流节点
func (f *FuncNode) To(funcNode Func) Node {
	node := NewFuncNode(funcNode)
	f.nextFuncFlows = node
	return node
}

// NewFuncNode 新建一个函数流节点
func NewFuncNode(funcNode Func) Node {
	return &FuncNode{funcNode: funcNode}
}

// ToNode 数据流入流节点
func (f *FuncNode) ToNode(node Node) Node {
	f.nextFuncFlows = node
	return node
}

// Next 下一个流节点
func (f *FuncNode) Next() Node {
	return f.nextFuncFlows
}

// Run 执行流节点函数
func (f *FuncNode) Run(in *Context) {
	f.funcNode(in)
	return
}
