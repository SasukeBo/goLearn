package main

// Printer doc false
type Printer interface {
	Print()
}

// S doc false
type S struct{}

// Print doc false
func (s S) Print() {
	println("print")
}

func main() {
	var i Printer

	// 没有初始化的接口调用其方法会产生panic
	// panic: runtime error: invalid memory address or nil pointer dereference
	// i.Print()

	// 必须初始化
	i = S{}
	i.Print()
}
