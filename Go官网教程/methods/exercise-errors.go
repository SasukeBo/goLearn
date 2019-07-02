package main

/*
	练习：错误

	Sqrt接收到一个负数时，应当返回一个非nil的错误值。复数同样也不被支持。
	创建一个新的类型

	type ErrNegativeSqrt float64

	并为其实现

	func (e ErrNegativeSqrt) Error() string

	方法使其拥有error值，通过ErrNegativeSqrt(-2).Error() 调用该方法应返回

	"cannot Sqrt negative number: -2"

	注意：在 Error 方法内调用fmt.Sprint(e) 会让程序陷入死循环。
	可以通过先转换e来避免这个问题: fmt.Sprint(float64(e))。

	修改Sqrt函数，使其接受一个负数时，返回ErrNegativeSqrt值。
*/
import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		e := ErrNegativeSqrt(x)
		return 0, e
	}
	return math.Sqrt(x), nil
}

func main() {
	if result, err := Sqrt(2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	if result, err := Sqrt(-2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
