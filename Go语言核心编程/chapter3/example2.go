package main

// X x
type X struct {
	a int
}

// Y y
type Y struct {
	X
}

// Z z
type Z struct {
	*X
}

// Get g
func (x X) Get() int {
	return x.a
}

// Set s
func (x *X) Set(i int) {
	x.a = i
}

func main() {
	x := X{a: 1}

	y := Y{
		X: x,
	}

	println(y.Get())

	// 此处编译器做了自动转换

	y.Set(2)
	println(y.Get())

	// 为了不让编译器做自动转换，使用方法表达式调用方式
	// Y内嵌字段X，所以type Y的方法集是Get，type *Y的方法集是Set Get
	(*Y).Set(&y, 3)

	// type Y 的方法集合并没有Set方法，所以下一句编译不能通过
	// Y.Set(y, 3)

	println(y.Get())

	z := Z{
		X: &x,
	}

	// 按照嵌套字段的方法集的规则
	// Z内嵌字段*X，所以type Z和type *Z方法集都包含类型X定义的方法Get和Set

	// 为了不让编译器做自动转换，仍然使用方法表达式调用方式
	Z.Set(z, 4)
	println(z.Get())

	(*Z).Set(&z, 5)
	println(z.Get())
}
