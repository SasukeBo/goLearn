package main

/*
基本的 for 循环由三部分组成，它们用分号隔开：

	初始化语句：在第一次迭代前执行。
	条件表达式：在每次迭代前求值。
	后置语句：在每次迭代的结尾执行。
*/
import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	sum2 := 1

	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)
}
