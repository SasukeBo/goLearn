package main

import "fmt"

// User doc false
// 使用type自定义的结构类型属于命名类型
type User struct {
	name string
	age  int
}

// errorString 是一个自定义结构类型，也是命名类型
type errorString struct {
	s string
}

func main() {
	// 结构字面量属于未命名类型
	a := struct {
		a int
		b string
	}{a: 1, b: "ok"}

	// struct{} 是非命名类型空结构
	var b = struct{}{}
	c := User{name: "sasuke", age: 25}
	d := errorString{s: "something goes wrong!"}

	fmt.Printf("a is type of %T, value is: %v\n", a, a)
	fmt.Printf("b is type of %T, value is: %v\n", b, b)
	fmt.Printf("c is type of %T, value is: %v\n", c, c)
	fmt.Printf("d is type of %T, value is: %v\n", d, d)
}
