package main

import "fmt"

// Student is a string key value map
type Student map[string]string

// PrintKeys println map keys
func (m Student) PrintKeys() {
	//底层类型支持的range运算，新类型可用
	for key, value := range m {
		fmt.Println(key, value)
	}
}

// MyInt is my int type
type MyInt int

func main() {
	var a MyInt = 10
	var b MyInt = 10

	// int 类型支持的加减乘除运算，新类型同样可用
	c := a + b
	d := a * b

	fmt.Printf("c is: %d\n", c)
	fmt.Printf("d is: %d\n", d)

	sasuke := Student{"name": "sasuke", "sex": "male", "telephone": "13242931765"}
	sasuke.PrintKeys()
}
