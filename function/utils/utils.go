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

package utils

import (
	"github.com/yunqi/flow"
	"os"
)

func ToStrings(in flow.Data) (strs []string) {
	switch in.(type) {
	case string:
		strs = append(strs, in.(string))
	case []string:
		strs = append(strs, in.([]string)...)
	}
	return
}
func ToFiles(in flow.Data) (files []*os.File) {
	switch in.(type) {
	case *os.File:
		files = append(files, in.(*os.File))
	case []*os.File:
		files = append(files, in.([]*os.File)...)
	}
	return
}
