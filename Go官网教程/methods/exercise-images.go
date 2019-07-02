package main

/*
练习：图像

定义自己的 Image 类型，实现必要的方法并调用 pic.ShowImage。
Bounds 应当返回一个 image.Rectangle，例如 image.Rect(0, 0, w, h)。
ColorModel 应当返回 color.RGBAModel。
At 应当返回一个颜色。上一个图片生成器的值 v 对应于此次的 color.RGBA{v, v, 255, 255}。
*/

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	width  int
	height int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
	img_func := func(x, y int) uint8 {
		return uint8(x ^ y)
	}
	v := img_func(x, y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{20, 300}
	pic.ShowImage(m)
}
