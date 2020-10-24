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
	Next() Node               // 子计算流
	Run(in *Data) (out *Data) //
	To(funcNode Func) Node
	ToNode(node Node) Node
}

// Data 数据流
type Data struct {
	flowId string
	data   interface{}
	err    error
}

func (d *Data) String() string {
	if d.err != nil {
		return fmt.Sprintf("{ flowId:%s,data:%v, err:%v}", d.flowId, d.data, d.err)
	} else {
		return fmt.Sprintf("{ flowId:%s,data:%v}", d.flowId, d.data)
	}
}

func NewData(data interface{}) *Data {
	flowId := strings.ReplaceAll(uuid.Must(uuid.NewV4(), nil).String(), "-", "")
	return &Data{
		data:   data,
		flowId: flowId,
	}
}
func (d *Data) Get() interface{} {
	return d.data
}
func (d *Data) Set(data interface{}) {
	d.data = data
}
func (d *Data) Err() error {
	return d.err
}
func (d *Data) SetErr(err error) {
	d.err = err
}
