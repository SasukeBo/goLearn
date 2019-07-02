/*
切片

实现Pic。他应当返回一个长度为 dy 的切片，其中每个元素是一个长度为 dx，元素类型为 uint8
的切片。当运行此程序时，它会将每个整数解释为灰度值并显示它所对应的图像。

图像的选择由你来定。几个有趣的函数包括 (x+y)/2, x*y, x^y, x*log(y) 和 x%(y+1)。

（提示：需要使用循环来分配[][]uint8中的每个[]uint8；请使用uint8(intValue)在类型之间
转换；可能会用到math包中的函数。）
*/
package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		imageI := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			imageI[j] = uint8(i * j)
		}
		image[i] = imageI
	}

	return image
}

func main() {
	pic.Show(Pic)
}
