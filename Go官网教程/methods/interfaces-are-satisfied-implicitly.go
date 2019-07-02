package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// 此方法表示类型T实现了接口I，但我们无需显式声明此事。
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"Hello"}
	i.M()
}

/*
类型通过实现一个接口的所有方法来实现接口。无需专门显式声明。

隐式接口从接口的实现中解耦了定义，这样接口的实现可以出现在任何包中，无需提前准备。
*/
