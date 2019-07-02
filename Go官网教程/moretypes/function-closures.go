/*
Go 函数可以是一个闭包。闭包是一个函数值，他引用了其函数体之外的变量。该函数可以访问并
赋予其引用的变量的值，换句话说，该函数被这些变量绑定在一起。
例如，函数adder返回一个闭包。每个闭包都被绑定在其各自的sum变量上。
*/
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}
