package main

/*
也可以为非结构体类型声明方法。
只能为同一包内定义的类型的接收者声明方法。
*/

import (
	"fmt"
	"math"
)

// MyFloat 数值类型
type MyFloat float64

// Abs 的接收者是MyFloat
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
