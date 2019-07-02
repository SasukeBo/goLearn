package main

import "fmt"

// T doc false
type T struct {
	a int
}

// Get a of T struct instance
func (t T) Get() int {
	return t.a
}

// Set a of T struct instance
func (t *T) Set(i int) {
	t.a = i
}

func main() {
	var t = &T{}
	t.Set(2)

	fmt.Println("t.Get() =>", t.Get())
	fmt.Println("t.Get =>", t.Get)
	f := t.Get
	fmt.Printf("f type is: %T\n", f)
	fmt.Println("f() => ", f())
}
