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
	"os"
	"strconv"
	"sync/atomic"
	"testing"
	"time"
)

func TestFlowBuffer(t *testing.T) {

	var buffer bytes.Buffer
	rand.Seed(2020)
	i := 0
	flow := NewFlow(1)

	flow2 := flow.To(func(in Data) (Data, bool) {
		b := in.(*bytes.Buffer)
		data, err := ioutil.ReadAll(b)
		var buffer bytes.Buffer

		if err != nil {
			fmt.Println("错误")

			return &buffer, true
		} else {
			var buffer bytes.Buffer

			d := string(data) + strconv.Itoa(i) + "node1"
			buffer.Write([]byte(d))
			i++

			return &buffer, true
		}

	})
	flow2.To(func(in Data) (Data, bool) {
		b := in.(*bytes.Buffer)

		time.Sleep(2 * time.Millisecond)
		data, err := ioutil.ReadAll(b)
		var buffer bytes.Buffer
		if err != nil {
			fmt.Println("错误")
			return &buffer, true
		} else {

			d := string(data) + "node2\n"
			buffer.Write([]byte(d))

		}

		return &buffer, true
	})

	flow.Run(true)
	f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0664)
	if err != nil {
		return
	}
	var j int64 = 0
	for i := 0; i < 10000; i++ {
		flow.Feed(&buffer, func(data Data) {
			b := data.(*bytes.Buffer)

			dataBytes, err := ioutil.ReadAll(b)
			if err == nil {
				f.Write(dataBytes)
				fmt.Println(string(dataBytes))
			}
			atomic.AddInt64(&j, 1)
		})
	}
	flow.Wait()
	fmt.Println("j:", j)
}
func TestFlowNumber(t *testing.T) {

	flow := NewFlow(20)
	flow1 := flow.To(func(in Data) (Data, bool) {
		b := in.(int)
		return (rand.Intn(1000)) + b, true
	})
	flow1.To(func(in Data) (Data, bool) {
		b := in.(int)
		return (rand.Intn(1000)) + b, true
	})
	flow.Run(true)

	for i := 0; i < 1000; i++ {
		flow.Feed(rand.Intn(100), func(data Data) {
			fmt.Println(data)
		})
	}
	flow.Wait()
}
