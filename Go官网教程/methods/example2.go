package main

import (
	"fmt"
	"math"
)

// Vertex doc false
type Vertex struct {
	X, Y float64
}

// Abs 现在这个Abs的写法就是个正常的函数，功能并没有变化
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
