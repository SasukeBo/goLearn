package main

import "fmt"

// Map doc false
type Map map[string]string

// Print doc false
func (m Map) Print() {
	for _, key := range m {
		fmt.Println(key)
	}
}

type iMap Map

// 只要底层类型是slice、map等支持range的类型字面量，新类型仍然可以使用range迭代

// Print doc false
func (m iMap) Print() {
	for _, key := range m {
		fmt.Println(key)
	}
}

type slice []int

func (s slice) Print() {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	mp := make(map[string]string, 10)
	mp["hi"] = "tata"

	// mp 与 ma 有相同的底层类型 map[string]string, 并且mp是未命名类型变量
	// 所以 mp 可以直接赋值给 ma
	var ma Map = mp

	// 不能赋值，如下语句不能通过编译：
	// var im iMap = ma
	// im.Print()
	// im 与 ma 虽然有相同的底层类型 map[string]string, 但它们中没有一个是未命名类型

	ma.Print()

	// Map 实现了 Print()，所以其可以赋值给接口类型变量
	var i interface {
		Print()
	} = ma

	i.Print()

	s1 := []int{1, 2, 3}
	var s2 slice
	s2 = s1
	s2.Print()
}
