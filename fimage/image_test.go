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
	"image"
	"testing"
)

func TestResize(t *testing.T) {
	testCases := []struct {
		name string
		src  image.Image
		w, h int
		f    imaging.ResampleFilter
		want *image.NRGBA
	}{
		{
			"Resize 2x2 1x1 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			1, 1,
			imaging.Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1 * 4,
				Pix:    []uint8{0x55, 0x55, 0x55, 0xc0},
			},
		},
		{
			"Resize 2x2 1x2 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			1, 2,
			imaging.Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 1 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0x80,
					0x00, 0x80, 0x80, 0xff,
				},
			},
		},
		{
			"Resize 2x2 2x1 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			2, 1,
			imaging.Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 2, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0xff, 0x00, 0x80, 0x80, 0x00, 0x80, 0xff,
				},
			},
		},
		{
			"Resize 2x2 2x2 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			2, 2,
			imaging.Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
		},
		{
			"Resize 3x1 1x1 nearest",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 2, 0),
				Stride: 3 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			1, 1,
			imaging.NearestNeighbor,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1 * 4,
				Pix:    []uint8{0x00, 0xff, 0x00, 0xff},
			},
		},
		{
			"Resize 2x2 0x4 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			0, 4,
			imaging.Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 4, 4),
				Stride: 4 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
		},
		{
			"Resize 2x2 4x0 linear",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			4, 0,
			imaging.Linear,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 4, 4),
				Stride: 4 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x40, 0xff, 0x00, 0x00, 0xbf, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0x40, 0x6e, 0x6d, 0x25, 0x70, 0xb0, 0x14, 0x3b, 0xcf, 0xbf, 0x00, 0x40, 0xff,
					0x00, 0xff, 0x00, 0xbf, 0x14, 0xb0, 0x3b, 0xcf, 0x33, 0x33, 0x99, 0xef, 0x40, 0x00, 0xbf, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0xbf, 0x40, 0xff, 0x00, 0x40, 0xbf, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
		},
		{
			"Resize 0x0 1x1 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, -1, -1),
				Stride: 0,
				Pix:    []uint8{},
			},
			1, 1,
			imaging.Box,
			&image.NRGBA{},
		},
		{
			"Resize 2x2 0x0 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			0, 0,
			imaging.Box,
			&image.NRGBA{},
		},
		{
			"Resize 2x2 -1x0 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			-1, 0,
			imaging.Box,
			&image.NRGBA{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//tc.src,

			got := Resize(tc.w, tc.h, tc.f)
			ims, _ := got(tc.src)
			newIms := ims.([]*image.NRGBA)
			for _, im := range newIms {
				if !compareNRGBA(im, tc.want, 0) {
					t.Fatalf("got result %#v want %#v", got, tc.want)
				}
			}

		})
	}
}

func TestResampleFilters(t *testing.T) {
	for _, filter := range []imaging.ResampleFilter{
		imaging.NearestNeighbor,
		imaging.Box,
		imaging.Linear,
		imaging.Hermite,
		imaging.MitchellNetravali,
		imaging.CatmullRom,
		imaging.BSpline,
		imaging.Gaussian,
		imaging.Lanczos,
		imaging.Hann,
		imaging.Hamming,
		imaging.Blackman,
		imaging.Bartlett,
		imaging.Welch,
		imaging.Cosine,
	} {
		t.Run("", func(t *testing.T) {
			src := image.NewNRGBA(image.Rect(-1, -1, 2, 3))
			got := Resize(5, 6, filter)
			ims, _ := got(src)
			newIms := ims.([]*image.NRGBA)

			want := image.NewNRGBA(image.Rect(0, 0, 5, 6))
			if !compareNRGBA(newIms[0], want, 0) {
				t.Fatalf("got result %#v want %#v", got, want)
			}
			if filter.Kernel != nil {
				if x := filter.Kernel(filter.Support + 0.0001); x != 0 {
					t.Fatalf("got kernel value %f want 0", x)
				}
			}
		})
	}
}

func compareNRGBA(img1, img2 *image.NRGBA, delta int) bool {
	if !img1.Rect.Eq(img2.Rect) {
		return false
	}
	return compareBytes(img1.Pix, img2.Pix, delta)
}
func compareBytes(a, b []uint8, delta int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if absint(int(a[i])-int(b[i])) > delta {
			return false
		}
	}
	return true
}
func absint(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
