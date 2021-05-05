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
	"math/rand"
	"testing"
)

func TestFlowNumber(t *testing.T) {

	flow := NewFlow(20)
	flow1 := flow.To(func(in *Context) {
		b := in.Data().(int)
		in.SetData((rand.Intn(1000)) + b)
	})
	flow1.To(func(in *Context) {
		b := in.Data().(int)
		in.SetData((rand.Intn(1000)) + b)
	})
	flow.Run(true)

	for i := 0; i < 1000; i++ {
		func(n int) {
			feedId := flow.Feed(n, func(data *Context) {
				fmt.Println(data)
			})
			fmt.Println(feedId)
		}(rand.Intn(100))

	}

	flow.Wait()
}
