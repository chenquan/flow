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

package file

import (
	"fmt"
	"github.com/yunqi/flow"
	"github.com/yunqi/flow/function/ffile"
	"os"
	"testing"
)

func TestImage(t *testing.T) {
	newFlow := flow.NewFlow(10)
	newFlow.To(ffile.GetAllFiles(".jpg")).
		To(ffile.OpenFile(os.O_RDWR, 0664)).
		To(ffile.GetSize())

	newFlow.Run(false)
	paths := []string{"data/", "d", "11/", "../"}

	for _, path := range paths {

		newFlow.Feed(path, func(result *flow.Context) {
			fmt.Println(result)
		})
	}

	newFlow.Wait()
}
