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
	"github.com/disintegration/imaging"
	"github.com/yuanqi/flow"
	"image"
)

// OpenWithPath 打开从流中获取的图片
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
			return ims, len(ims) != 0
		}
		return nil, false
	}
}

// CropAnchor 裁剪
func CropAnchor(width, height int, anchor imaging.Anchor) flow.Func {
	return func(in flow.Data) (flow.Data, bool) {
		var ims []image.Image
		for _, im := range toImages(in) {
			newIm := imaging.CropAnchor(im, width, height, anchor)
			ims = append(ims, newIm)
		}
		return ims, len(ims) != 0

	}
}

// toImages 转为 []image.Image
func toImages(in flow.Data) (images []image.Image) {
	switch in.(type) {
	case image.Image:
		images = append(images, in.(image.Image))
	case []image.Image:
		images = append(images, in.([]image.Image)...)
	}
	return
}

// handleImages 使用处理函数去处理 image.Image
func handleImages(in flow.Data, handle func(im image.Image) *image.NRGBA) (ims []image.Image) {
	for _, im := range toImages(in) {
		newIm := handle(im)
		ims = append(ims, newIm)
	}
	return ims
}

// Resize 调整图片大小
func Resize(width, height int, filter imaging.ResampleFilter) flow.Func {
	return func(in flow.Data) (flow.Data, bool) {
		ims := handleImages(in, func(im image.Image) *image.NRGBA {
			return imaging.Resize(im, width, height, filter)
		})
		return ims, len(ims) != 0
	}
}

// Fit 按比例缩小图像使用指定的重采样滤波器，以适应指定的最大宽度和高度
func Fit(width, height int, filter imaging.ResampleFilter) flow.Func {
	return func(in flow.Data) (flow.Data, bool) {
		ims := handleImages(in, func(im image.Image) *image.NRGBA {
			return imaging.Fit(im, width, height, filter)
		})
		return ims, len(ims) != 0
	}
}

// Fill 调整并裁剪图片
func Fill(width, height int, anchor imaging.Anchor, filter imaging.ResampleFilter) flow.Func {
	return func(in flow.Data) (flow.Data, bool) {
		ims := handleImages(in, func(im image.Image) *image.NRGBA {
			return imaging.Fill(im, width, height, anchor, filter)
		})

		return ims, len(ims) != 0
	}
}

// Sharpen 会生成图像的锐化版本
func Sharpen(sigma float64) flow.Func {
	return func(in flow.Data) (flow.Data, bool) {
		ims := handleImages(in, func(im image.Image) *image.NRGBA {
			return imaging.Sharpen(im, sigma)
		})
		return ims, len(ims) != 0
	}
}

// AdjustGamma 对图像执行gamma校正,然后返回调整后的图像
func AdjustGamma(gamma float64) flow.Func {
	return func(in flow.Data) (flow.Data, bool) {
		ims := handleImages(in, func(im image.Image) *image.NRGBA {
			return imaging.AdjustGamma(im, gamma)
		})
		return ims, len(ims) != 0
	}
}

// AdjustContrast 使用percent参数更改图像的对比度,并返回调整后的图像
func AdjustContrast(percentage float64) flow.Func {
	return func(in flow.Data) (flow.Data, bool) {
		ims := handleImages(in, func(im image.Image) *image.NRGBA {
			return imaging.AdjustContrast(im, percentage)
		})
		return ims, len(ims) != 0
	}
}

// AdjustBrightness 使用percentage参数更改图像的亮度,并返回调整后的图像
func AdjustBrightness(percentage float64) flow.Func {
	return func(in flow.Data) (flow.Data, bool) {
		ims := handleImages(in, func(im image.Image) *image.NRGBA {
			return imaging.AdjustBrightness(im, percentage)
		})
		return ims, len(ims) != 0
	}
}

// AdjustSaturation 使用percentage参数更改图像的饱和度,并返回调整后的图像
func AdjustSaturation(percentage float64) flow.Func {
	return func(in flow.Data) (flow.Data, bool) {
		ims := handleImages(in, func(im image.Image) *image.NRGBA {
			return imaging.AdjustSaturation(im, percentage)
		})
		return ims, len(ims) != 0
	}
}

// Blur 使用高斯函数,生成图像的模糊版本
func Blur(sigma float64) flow.Func {
	return func(in flow.Data) (flow.Data, bool) {
		ims := handleImages(in, func(im image.Image) *image.NRGBA {
			return imaging.Blur(im, sigma)
		})
		return ims, len(ims) != 0
	}
}

// Invert 会生成图像的反转版本
func Invert() flow.Func {
	return func(in flow.Data) (flow.Data, bool) {
		ims := handleImages(in, func(im image.Image) *image.NRGBA {
			return imaging.Invert(im)
		})
		return ims, len(ims) != 0
	}
}

// Grayscale 产生图像的灰度版本
func Grayscale() flow.Func {
	return func(in flow.Data) (flow.Data, bool) {
		ims := handleImages(in, func(im image.Image) *image.NRGBA {
			return imaging.Grayscale(im)
		})
		return ims, len(ims) != 0
	}
}

// Convolve3x3 使用指定的3x3卷积内核对图像进行卷积
func Convolve3x3(kernel [9]float64, options *imaging.ConvolveOptions) flow.Func {

	return func(in flow.Data) (flow.Data, bool) {
		ims := handleImages(in, func(im image.Image) *image.NRGBA {
			return imaging.Convolve3x3(
				im,
				kernel,
				options,
			)
		})
		return ims, len(ims) != 0
	}
}

// Paste 将img图像粘贴到指定位置的背景图像，然后返回合并的图像
func Paste(img image.Image, pos image.Point) flow.Func {
	return func(in flow.Data) (flow.Data, bool) {
		ims := handleImages(in, func(im image.Image) *image.NRGBA {
			return imaging.Paste(im, img, pos)
		})
		return ims, len(ims) != 0
	}
}
