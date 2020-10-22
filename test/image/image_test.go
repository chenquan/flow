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

package image

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/yuanqi/flow"
	fimgae "github.com/yuanqi/flow/fimage"
	"image"
	"math/rand"
	"strconv"
	"testing"
)

func TestImage(t *testing.T) {
	newFlow := flow.NewFlow(10)
	pathFlow := newFlow.FlowIn(fimgae.GetAllFiles(".jpg"))
	openFlow := pathFlow.FlowIn(fimgae.OpenWithPath())
	h1 := openFlow.FlowIn(fimgae.CropAnchor(300, 300, imaging.Center))
	h1.FlowIn(fimgae.CropAnchor(100, 300, imaging.Bottom))

	newFlow.Run()
	paths := []string{"data/", "a/"}

	rand.Seed(2020)
	for _, path := range paths {
		newFlow.Feed(path, func(result flow.Data) {
			if ims, ok := result.([]image.Image); ok {
				fmt.Println(ims)
				for _, im := range ims {
					_ = imaging.Save(im, strconv.Itoa(rand.Int())+".jpg")
				}
			}
		})
	}

	newFlow.Wait()
}
