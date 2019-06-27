package main

import (
	"fmt"
)

func test(i int) (j int) {
	return i
}

func main() {
	fmt.Printf("%T\n", test)
}
