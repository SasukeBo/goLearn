package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	// 以值为接收者的方法被调用时，接受者既能为值也能为指针
	p := &Vertex{4, 3}
	// 这种情况下，方法调用 p.Abs() 会被解释为 (*p).Abs()。
	fmt.Println(p.Abs())
	fmt.Println((*p).Abs())
	fmt.Println(AbsFunc(*p))
}
