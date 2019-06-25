package main

import (
	"fmt"
)

func main() {
	var a [2]int
	s := a[0:2]

	fmt.Printf("s is %v\n", s)

	l := append(s, 0)
	printSlice(l)
	fmt.Printf("s is %v\n", s)

	l = append(l, 1)
	printSlice(l)
	fmt.Printf("s is %v\n", s)

	l = append(l, 2, 3, 4)
	printSlice(l)
	fmt.Printf("s is %v\n", s)
}

func printSlice(s []int) {
	fmt.Printf("len = %d cap = %d %v\n", len(s), cap(s), s)
}
