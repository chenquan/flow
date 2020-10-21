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
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"sync/atomic"
	"testing"
)

func TestFlow(t *testing.T) {

	var buffer bytes.Buffer
	rand.Seed(2020)
	//input := NewInput(&buffer)
	i := 0
	node1 := NewFuncNode(func(in Data) Data {
		b := in.(*bytes.Buffer)
		data, err := ioutil.ReadAll(b)
		var buffer bytes.Buffer
		//time.Sleep(time.Duration(rand.Intn(1)) * time.Second)

		if err != nil {
			fmt.Println("错误")

			return &buffer
		} else {
			var buffer bytes.Buffer

			d := string(data) + strconv.Itoa(i) + "node1"
			buffer.Write([]byte(d))
			i++

			return &buffer
		}

	})
	node2 := NewFuncNode(func(in Data) Data {
		b := in.(*bytes.Buffer)

		//time.Sleep(time.Duration(rand.Intn(2)) * time.Millisecond)
		data, err := ioutil.ReadAll(b)
		var buffer bytes.Buffer
		if err != nil {
			fmt.Println("错误")
			return &buffer
		} else {

			d := string(data) + "node2\n"
			buffer.Write([]byte(d))

		}

		return &buffer
	})

	flow := NewFlow(1)
	flow2 := flow.FlowIn(node1)
	flow2.FlowIn(node2)
	flow.Run()

	var j int64 = 0
	for i := 0; i < 10000; i++ {
		flow.Feed(&buffer, func(data Data) {
			b := data.(*bytes.Buffer)

			dataBytes, err := ioutil.ReadAll(b)
			if err == nil {
				fmt.Println(string(dataBytes))
			}
			atomic.AddInt64(&j, 1)
		})
	}
	flow.Wait()
	fmt.Println("j:", j)
}
