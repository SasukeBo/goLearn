package main

import "fmt"

func main() {
	s := "hello, 世界！"
	var a []byte
	a = ([]byte)(s)
	var b string
	b = (string)(a)

	var c []rune
	c = ([]rune)(s)

	fmt.Printf("%T\n", a) // []uint8 - byte 是 int8 的别名
	fmt.Printf("%T\n", b) // string
	fmt.Printf("%T\n", c) // []uint32 - rune 是 int32 的别名

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
}
