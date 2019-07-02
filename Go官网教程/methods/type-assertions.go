package main

/*
类型断言提供了访问接口值底层具体值的方式

t := i.(T)

该语句断言接口值i保存了具体类型 T，并将其底层类型为 T 的值赋予变量 t。
若 i 并未保存 T 类型的值，该语句就会触发一个 panic。

为了判断一个接口值是否保存了一个特定的类型，类型断言可返回两个值：其底层值以及一个报告
断言是否成功的布尔值。

若 i 保存了一个 T，那么 t 将会是其底层值，而 ok 为true。
否则，ok将会为false，而t将为T类型零值，程序不会产生panic。


*/

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64) // 报错
	// fmt.Println(f)
}
