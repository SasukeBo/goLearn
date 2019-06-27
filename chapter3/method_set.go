package main

import "fmt"

// Int doc false
type Int int

// Max doc false
func (i Int) Max(b Int) Int {
	if i >= b {
		return i
	}

	return b
}

// Set doc false
func (i *Int) Set(a Int) {
	*i = a
}

// Print doc false
func (i Int) Print() {
	fmt.Printf("Value=%d\n", i)
}

func main() {
	var a Int = 10
	var b Int = 20

	c := a.Max(b)
	c.Print()
	(&c).Print()

	a.Set(20)
	a.Print()

	(&a).Set(30)
	a.Print()
}
