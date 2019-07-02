package main

import "fmt"

// Person struct named type
// 使用type声明的是命名类型
type Person struct {
	name string
	age  int
}

func main() {
	// 使用struct字面量声明的是未命名类型
	a := struct {
		name string
		age  int
	}{"andes", 18}

	fmt.Printf("%T\n", a) // struct { name string; age int }
	fmt.Printf("%v\n", a) // { andes 18 }

	b := Person{"sasuke", 25}
	fmt.Printf("%T\n", b) // main.Person
	fmt.Printf("%v\n", b) // { sasuke, 25 }

	c := [10]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", c)
	fmt.Printf("%v\n", c)
}
