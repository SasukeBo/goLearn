package main

import "fmt"

// X doc false
type X struct {
	a int
}

// Y doc false
type Y struct {
	X
	b int
}

// Z doc false
type Z struct {
	Y
	c int
}

// Print doc false
func (x X) Print() {
	fmt.Printf("In X, a = %d\n", x.a)
}

// XPrint doc false
func (x X) XPrint() {
	fmt.Printf("In X, a = %d\n", x.a)
}

// Print doc false
func (y Y) Print() {
	fmt.Printf("In Y, b = %d\n", y.b)
}

// Print doc false
func (z Z) Print() {
	fmt.Printf("In Z, c = %d\n", z.c)
	// 显示的完全路径调用内嵌字段的方法
	z.Y.Print()
	z.Y.X.Print()
}

func main() {
	x := X{a: 1}

	y := Y{
		X: x,
		b: 2,
	}

	z := Z{
		Y: y,
		c: 3,
	}

	// 从外向内查找，首先查找到的是Z的Print()方法
	z.Print()
	// 从外向内查找，最后找到的是X的XPrint()方法
	z.XPrint()
	z.Y.XPrint()
}
