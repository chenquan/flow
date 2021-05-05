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
	"testing"
	"time"
)

func TestFlowNumber(t *testing.T) {

	flow := NewFlow(20)
	flow1 := flow.To(func(in *Context) {
		b := in.Data().(int)
		in.SetData(1 + b)
	})
	flow1.To(func(in *Context) {
		b := in.Data().(int)
		in.SetData((3) + b)
	})
	flow.Run(func(data *Context) {
		fmt.Println(data)
	}, WithPoolSize(20000))

	for i := 0; i < 1000; i++ {
		func(n int) {
			_ = flow.Feed(n)
			//fmt.Println(feedId)
		}(1)

	}

	flow.Wait()
}

func BenchmarkFlowNumberWithPool(b *testing.B) {

	flow := NewFlow(20)
	flow1 := flow.To(func(in *Context) {
		b := in.Data().(int)
		in.SetData(1 + b)
		time.Sleep(time.Microsecond)
	})
	flow1.To(func(in *Context) {
		b := in.Data().(int)
		in.SetData((3) + b)
	})
	flow.Run(func(data *Context) {
		//fmt.Println(data)
	})
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = flow.Feed(1)
		}
	})
	flow.Wait()

}
func BenchmarkFlowNumberWithOutPool(b *testing.B) {

	flow := NewFlow(20)
	flow1 := flow.To(func(in *Context) {
		b := in.Data().(int)
		in.SetData(1 + b)
		time.Sleep(time.Microsecond)
	})
	flow1.To(func(in *Context) {
		b := in.Data().(int)
		in.SetData((3) + b)
	})
	flow.Run(func(data *Context) {
		//fmt.Println(data)
	}, WithDisablePool())
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = flow.Feed(1)
		}
	})
	flow.Wait()

}
