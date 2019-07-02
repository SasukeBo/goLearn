package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6}
	b := make([]int, 2, 4)
	c := a[0:3]

	fmt.Println(len(b))
	fmt.Println(cap(b))
	b = append(b, 1)
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))

	b = append(b, c...)
	fmt.Println(b)
	fmt.Println("slice b len is", len(b))
	fmt.Println("slice b cap is", cap(b))
	fmt.Println("slice c len is:", len(c))
	fmt.Println("slice c cap is:", cap(c))

	d := make([]int, 2, 2)
	copy(d, c)
	fmt.Println(d)
	fmt.Println(len(d))
	fmt.Println(cap(d))
}
