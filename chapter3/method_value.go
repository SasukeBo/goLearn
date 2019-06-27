package main

import "fmt"

// Data contain int length
type Data struct {
	length int
}

// GetLen return int length value
func (d Data) GetLen() int {
	return d.length
}

// SetLen set data length
func (d *Data) SetLen(len int) {
	d.length = len
}

// Print doc false
func (d Data) Print() {
	fmt.Printf("length is: %v\n", d.GetLen())
}

func main() {
	// 声明一个类型变量a
	var a Data = struct{ length int }{length: 100}

	// 表达式调用编译器不会进行自动转换
	a.Print()
	// Data.GetLen(&a) // error
	(*Data).SetLen(&a, 200)
	fmt.Printf("a.length is: %v\n", Data.GetLen(a))

	a.SetLen(300)
	f := a.GetLen
	fmt.Printf("a.length is: %v\n", f())

	a.SetLen(400)
	y := (&a).GetLen
	fmt.Printf("a.length is: %v\n", y())

	g := a.SetLen
	g(700)
	a.Print()

	x := (&a).SetLen
	x(800)
	a.Print()
}
