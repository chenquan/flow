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

package ffile

import (
	"fmt"
	"github.com/yunqi/flow"
	"github.com/yunqi/flow/function/utils"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func getAllFiles(dirPath string, suffix string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPath)

	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	suffix = strings.ToLower(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {

		if fi.IsDir() {
			dirs = append(dirs, dirPath+PthSep+fi.Name())
			paths, _ := getAllFiles(dirPath+PthSep+fi.Name(), suffix)
			files = append(files, paths...)
		} else {
			ext := path.Ext(fi.Name())
			if ext == suffix {
				files = append(files, dirPath+PthSep+fi.Name())
			}
		}

	}
	return files, nil
}

// GetAllFiles 获取指定后缀名的文件路径
func GetAllFiles(suffix string) flow.Func {

	return func(ctx *flow.Context) {
		if dirPath, ok := ctx.Data().(string); ok {
			files, err := getAllFiles(dirPath, suffix)
			if err == nil {
				ctx.SetData(files)
			} else {
				ctx.SetErr(err)
			}

		} else {
			ctx.SetErr(flow.Error)
		}
	}
}

// OpenFile 打开文件
func OpenFile(flag int, perm os.FileMode) flow.Func {
	return func(ctx *flow.Context) {
		var files []*os.File
		for _, s := range utils.ToStrings(ctx) {
			f, err := os.OpenFile(s, flag, perm)
			if err == nil {
				files = append(files, f)
			}
		}
		if len(files) != 0 {
			ctx.SetData(files)
		} else {
			ctx.SetErr(flow.Error)
		}
	}
}

// MkDir 创建文件夹
func MkDir(perm os.FileMode) flow.Func {
	return func(in *flow.Context) {
		var dirs []string
		for _, s := range utils.ToStrings(in) {
			err := os.MkdirAll(s, perm)
			if err == nil {
				dirs = append(dirs, s)
			}
		}
		if len(dirs) != 0 {
			in.SetData(dirs)
		} else {
			in.SetErr(flow.Error)
		}
	}
}

type FileSize struct {
	Filename string
	Size     int
}

func (f *FileSize) String() string {
	return fmt.Sprintf("{Filename:%s, Size:%d}", f.Filename, f.Size)
}

// GetSize 获取文件大小
func GetSize() flow.Func {
	return func(ctx *flow.Context) {

		fileSizes := make([]*FileSize, 0)
		for _, file := range utils.ToFiles(ctx) {
			file.Name()
			size, err := ioutil.ReadAll(file)
			if err == nil {
				fileSizes = append(fileSizes, &FileSize{
					Filename: file.Name(),
					Size:     len(size),
				})
			}
		}
		if len(fileSizes) != 0 {
			ctx.SetData(fileSizes)
		} else {
			ctx.SetErr(flow.Error)
		}
	}
}

// GetExt 获取文件后辍名
func GetExt() flow.Func {
	return func(ctx *flow.Context) {
		var exts []string
		for _, fileName := range utils.ToStrings(ctx) {
			exts = append(exts, path.Ext(fileName))
		}
		if len(exts) != 0 {
			ctx.SetData(exts)
		} else {
			ctx.SetErr(flow.Error)
		}
	}
}
