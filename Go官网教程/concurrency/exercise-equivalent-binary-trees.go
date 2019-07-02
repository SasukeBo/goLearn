package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk doc false
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same doc false
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int, 10), make(chan int, 10)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	var ta1, ta2 [10]int

	for i := 0; i < 10; i++ {
		ta1[i] = <-ch1
		ta2[i] = <-ch2
	}

	return ta1 == ta2
}

func main() {
	ch := make(chan int, 10)
	t := tree.New(1)
	go Walk(t, ch)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
