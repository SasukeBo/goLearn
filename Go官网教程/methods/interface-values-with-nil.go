package main

// 底层值为nil的接口值
// 即便接口内的具体值为nil，方法仍然会被nil接收者调用
// 在一些语言中，这回触发一个空指针异常，但在Go中通常会写一些方法来优雅的处理它（如本例中的M方法）。
// 注意：保存了nil具体值的接口其自身并不为nil。
import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
