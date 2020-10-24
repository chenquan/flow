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
	"sync"
)

// Node 实现该接口的是计算流
type Node interface {
	Next() Node      // 子计算流
	Run(in *Context) //
	To(funcNode Func) Node
	ToNode(node Node) Node
}

// Context 流处理上下文
type Context struct {
	flowId string                 // 流ID
	data   interface{}            // 时间
	step   int32                  // 当前执行步骤
	err    error                  // 错误信息
	cache  map[string]interface{} // 缓存
	mu     sync.RWMutex           // 保护 cache
	once   sync.Once              // err 只能被更改一次
}

func (c *Context) String() string {
	if c.err != nil {
		return fmt.Sprintf("{ flowId:%s, step:%c, data:%v, err:%v}", c.flowId, c.step, c.data, c.err)
	} else {
		return fmt.Sprintf("{ flowId:%s, step:%c, data:%v}", c.flowId, c.step, c.data)
	}
}

func NewContext(data interface{}) *Context {
	flowId := strings.ReplaceAll(uuid.Must(uuid.NewV4(), nil).String(), "-", "")
	return &Context{
		data:   data,
		flowId: flowId,
		step:   -1,
		cache:  make(map[string]interface{}),
	}
}

// Data 返回数据
// 并发不安全
func (c *Context) Data() interface{} {
	return c.data
}

// SetData 修改数据
// 并发不安全
func (c *Context) SetData(data interface{}) {
	c.data = data
}

// Err 返回错误信息
func (c *Context) Err() error {
	return c.err
}

// SetErr 设置错误信息
// 只能被设置一次非nil错误信息
func (c *Context) SetErr(err error) {
	if err != nil {
		c.once.Do(func() {
			c.err = err
		})
	}

}

// FlowId 返回流处理ID
func (c *Context) FlowId() string {
	return c.flowId
}

// SetCache 设置缓存
func (c *Context) SetCache(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = value
}

// GetCache 返回对应 key 的缓存值
func (c *Context) GetCache(key string) (value interface{}) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.cache[key]
}
