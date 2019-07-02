package main

import "fmt"

// SliceInt doc false
type SliceInt []int

// Sum doc false
func (s SliceInt) Sum() int {
	sum := 0
	for _, i := range s {
		sum += i
	}

	return sum
}

// SliceIntSum 这个函数和上面的方法等价
func SliceIntSum(s SliceInt) int {
	sum := 0
	for _, i := range s {
		sum += i
	}

	return sum
}

func main() {
	var s SliceInt = []int{1, 2, 3, 4}
	fmt.Println(s.Sum())
	fmt.Println(SliceIntSum(s))
}
