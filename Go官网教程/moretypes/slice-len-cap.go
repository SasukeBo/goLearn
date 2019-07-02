package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 截取切片使其长度为0
	s = s[:0]
	printSlice(s)

	// 拓展其长度
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值
	// 切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数，
	// 改变切片的头指针就意味着改变其容量
	s = s[2:]
	printSlice(s)

	s = s[1:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
