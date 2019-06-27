package main

import "fmt"

func main() {
	type T1 string
	type T2 T1
	type T3 []string
	type T4 T3
	type T5 []T1
	type T6 T5

	var t1 T1 = "hello"
	var t2 T2 = "world"

	var t3 T3
	t3 = T3{"a", "b", "c"}
	t4 := T4{"a", "b", "c"}
	t5 := T5{"a", "b", "c"}
	t6 := T6{"a", "b", "c"}

	fmt.Printf("t1 value: %v, type: %T\n", t1, t1)
	fmt.Printf("t2 value: %v, type: %T\n", t2, t2)
	fmt.Printf("t3 value: %v, type: %T\n", t3, t3)
	fmt.Printf("t4 value: %v, type: %T\n", t4, t4)
	fmt.Printf("t5 value: %v, type: %T\n", t5, t5)
	fmt.Printf("t6 value: %v, type: %T\n", t6, t6)
}
