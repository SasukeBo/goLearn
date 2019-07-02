/*
练习：斐波那契闭包
实现一个fibonacci函数，它返回一个函数（闭包），该闭包返回一个斐波那契数列。
*/
package main

import "fmt"

// 返回一个“返回int的函数”

func fibonacci() func() int {
	num := []int{0, 1}

	return func() int {
		out := num[0]
		num[0] = num[1]
		num[1] = out + num[0]

		return out
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
