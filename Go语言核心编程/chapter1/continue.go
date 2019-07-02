package main

import "fmt"

func main() {
L1:
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			fmt.Println("i: ", i, "j: ", j)
			if i >= 10 {
				break L1
			}
			if i >= 5 {
				// 跳出L1标签所在的for循环i++处执行
				continue L1
				// the following is not executed
			}
			if j >= 100 {
				break
			}
			if j > 10 {
				// 默认仅跳到离break最近的内层循环j++处执行
				continue
			}
		}
	}
}
