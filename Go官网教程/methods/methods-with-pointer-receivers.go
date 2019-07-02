package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Abs 即使Abs并不需要修改接收者，这样可以节约调用方法是复制接收者的开销
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v. Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}

/*
使用指针接收者的原因有二:

首先，方法能够修改其接收者指向的值
其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构时，这样做会更加高效。
*/
