package main

import "fmt"

func main() {
L1:
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			fmt.Println("i: ", i, "j: ", j)
			if i >= 5 {
				// 跳出L1标签所在的for循环
				break L1
			}
			if j > 10 {
				// 默认仅跳出离break最近的内存循环
				break
			}
		}
	}
}
