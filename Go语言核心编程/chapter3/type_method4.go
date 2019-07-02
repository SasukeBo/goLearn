package main

import "fmt"

// T doc false
type T struct {
	a int
}

// Get doc false
func (t T) Get() int {
	return t.a
}

// Set doc false
func (t *T) Set(i int) {
	t.a = i
}

// Print doc false
func (t *T) Print() {
	fmt.Printf("%p, %d, %d \n", t, t, t.a)
}

func main() {
	/*
		var t = &T{}

		// method value
		f := t.Set

		// 方法值调用
		f(2)
		t.Print()

		// 方法值调用
		f(3)
		t.Print()
		fmt.Printf("%T\n", t.Set)
	*/

	t := T{a: 1}

	fmt.Println(t.Get())
	fmt.Println((T).Get(t))

	f1 := T.Get
	fmt.Println(f1(t))

	(*T).Set(&t, 2)
	f3 := (*T).Set
	f3(&t, 3)

	f2 := (T).Get
	fmt.Println(f2(t))
}
