package main

import "fmt"

// Inter doc false
type Inter interface {
	Ping()
	Pang()
}

// Anter doc false
type Anter interface {
	Inter
	String()
}

// St doc false
type St struct {
	Name string
}

// Ping doc false
func (St) Ping() {
	println("ping")
}

// Pang doc false
func (*St) Pang() {
	println("pang")
}

// String doc false
func (*St) String() {
	println("hello world")
}

func main() {
	st := &St{"sasuke"}
	var i interface{} = st
	// 判断i绑定的实例是否实现了接口类型Inter
	o := i.(Inter)
	o.Ping()
	o.Pang()

	// 如下语句会引发panic，因为i没有实现接口Anter
	p := i.(Anter)
	p.String()

	s := i.(*St)
	fmt.Printf("%s\n", s.Name)
}
