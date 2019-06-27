package main

import "fmt"

// 有名函数定义，函数名是add
// add 类型是函数字面量类型 func (int, int) int
func add(a, b int) int {
	return a + b
}

// 函数声明语句，用于Go代码调用汇编代码
// func add(int, int) int

// add 函数的签名，实际上就是 add 的字面量类型
// func (int, int) int

// 匿名函数不能独立存在，常作为函数参数、返回值，或者赋值给某个变量
// 匿名函数可以直接显式初始化
// 匿名函数的类型也是函数字面量类型 func (int, int) int
// func (a, b int) int {
// return a + b
// }

// 新定义函数类型ADD

// ADD 底层类型是函数字面量类型 func (int, int) int
type ADD func(int, int) int

// add 和 ADD 的底层类型相同，并且 add 是字面量类型
// 所以 add 可直接赋值给 ADD 类型的变量g
var g ADD = add

func main() {
	f := func(a, b int) int {
		return a + b
	}

	g(1, 2)
	f(1, 2)

	// f 和 add 的函数签名相同
	fmt.Printf("%T\n", f)   // func(int, int) int
	fmt.Printf("%T\n", add) // func(int, int) int
}
