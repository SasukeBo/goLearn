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

// 只要底层类型是 slice、map 等支持 range 的类型字面量，新类型仍然可以使用 range 迭代
func (m iMap) Print() {
	for _, key := range m {
		fmt.Println(key)
	}
}

func main() {
	mp := make(map[string]string, 10)
	mp["hi"] = "tata"
	// mp 与 ma 有相同的底层类型 map[string]string，并且mp是未命名类型
	var ma Map = mp
	ma.Print()

	// im 与 ma 虽然有相同的底层类型，但是二者中没有一个是字面量类型，不能直接赋值，
	// 可以强制进行类型转换
	// var im iMap = ma
	var im iMap
	im = (iMap)(ma)
	im.Print()

	im = mp
	im.Print()
}
