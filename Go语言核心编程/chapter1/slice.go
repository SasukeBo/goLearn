package main

import "fmt"

func main() {
	var array = [...]int{0, 1, 2, 3, 4, 5, 6}
	s1 := array[0:4]
	s2 := array[:4]
	s3 := array[2:]

	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", s2)
	fmt.Printf("%v\n", s3)
}
