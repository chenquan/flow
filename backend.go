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
	uuid "github.com/satori/go.uuid"
	"strings"
)

// Node 实现该接口的是计算流
type Node interface {
	Next() Node                     // 子计算流
	Run(in *Context) (out *Context) //
	To(funcNode Func) Node
	ToNode(node Node) Node
}

// Context 流处理上下文
type Context struct {
	flowId string
	data   interface{}
	step   int32
	err    error
}

func (d *Context) String() string {
	if d.err != nil {
		return fmt.Sprintf("{ flowId:%s, step:%d, data:%v, err:%v}", d.flowId, d.step, d.data, d.err)
	} else {
		return fmt.Sprintf("{ flowId:%s, step:%d, data:%v}", d.flowId, d.step, d.data)
	}
}

func NewData(data interface{}) *Context {
	flowId := strings.ReplaceAll(uuid.Must(uuid.NewV4(), nil).String(), "-", "")
	return &Context{
		data:   data,
		flowId: flowId,
		step:   -1,
	}
}
func (d *Context) Get() interface{} {
	return d.data
}
func (d *Context) Set(data interface{}) {
	d.data = data
}
func (d *Context) Err() error {
	return d.err
}
func (d *Context) SetErr(err error) {
	d.err = err
}
func (d *Context) FlowId() string {
	return d.flowId
}
