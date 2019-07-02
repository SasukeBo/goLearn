package main

import "fmt"

func main() {
	var a [2]string
	/*
	   数组的长度是其类型的一部分，因此数组不能改变大小
	*/
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
