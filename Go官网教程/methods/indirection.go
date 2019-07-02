package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2) // 即便v是个值而不是指针，带指针接收者的方法也能被直接调用。
	// 由于Scale方法有一个指针接收者，为方便起见，Go会将语句v.Scale(5)解释为(&v).Scale(5)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	// 以指针为接收者的方法被调用时，接收者既可以为值，也可以是指针。

	fmt.Println(v, p)
}
