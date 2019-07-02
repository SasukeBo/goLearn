package main

/*
	零值
	没有明确初始值的变量声明会被赋予它们的零值
	零值是：
		数值类型为0，
		布尔类型为false，
		字符串为“”（空字符串）。
*/

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
