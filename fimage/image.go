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

package fimage

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/yuanqi/flow"
	"image"
	"io/ioutil"
	"os"
	"path"
)

func getAllFiles(dirPath string, suffix string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPath)

	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {

		if fi.IsDir() {
			dirs = append(dirs, dirPath+PthSep+fi.Name())
			newfiles, _ := getAllFiles(dirPath+PthSep+fi.Name(), suffix)
			files = append(files, newfiles...)
		} else {
			ext := path.Ext(fi.Name())
			if ext == suffix {
				files = append(files, dirPath+PthSep+fi.Name())
			}
		}

	}
	return files, nil
}
func GetAllFiles(suffix string) flow.Func {

	return func(in flow.Data) (flow.Data, bool) {
		if dirPath, ok := in.(string); ok {
			files, err := getAllFiles(dirPath, suffix)
			if err == nil {
				fmt.Println(files)
				return files, true
			}
		}
		return nil, false
	}

}
func OpenWithPath() flow.Func {
	f := func(path string) (image.Image, error) {

		im, err := imaging.Open(path, imaging.AutoOrientation(true))
		if err == nil {
			return im, err
		}
		return nil, err
	}
	return func(in flow.Data) (flow.Data, bool) {
		switch in.(type) {
		case string:
			if im, err := f(in.(string)); err == nil {
				return im, true
			} else {
				return nil, false
			}
		case []string:
			paths := in.([]string)
			var ims []image.Image
			for _, pathName := range paths {
				if im, err := f(pathName); err == nil {
					ims = append(ims, im)
				}
			}
			if len(ims) == 0 {
				return nil, false
			} else {
				return ims, true
			}
		}
		return nil, false
	}
}
func CropAnchor(width, height int, anchor imaging.Anchor) flow.Func {

	return func(in flow.Data) (flow.Data, bool) {
		switch in.(type) {
		case image.Image:
			return imaging.CropAnchor(in.(image.Image), width, height, anchor), true
		case []image.Image:
			images := in.([]image.Image)
			var ims []image.Image
			for _, im := range images {
				cropIm := imaging.CropAnchor(im, width, height, anchor)
				ims = append(ims, cropIm)
			}
			if len(ims) == 0 {
				return nil, false
			} else {
				return ims, true
			}
		}

		return nil, false
	}
}
