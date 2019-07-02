package main

// 可以为指针接收者声明方法
// 这意味着对于某类型T，接收者的类型可以用*T的文法。
// T不能是像*int这样的指针。

import (
	"fmt"
	"math"
)

// Vertex a struct type name
type Vertex struct {
	X, Y float64
}

// Abs doc false
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale 这里为*Vertex定义了Scale方法。
// 指针接收者的方法可以修改接收者指向的值，由于方法经常需要修改
// 它的接收者，指针接收者比值接收者更常用。
// 若使用值接收者，那么Scale方法会对原始Vertex值的副本进行操
// 作。
// Scale方法必须用指针接收者来更改main函数中声明的Vertex的值
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
