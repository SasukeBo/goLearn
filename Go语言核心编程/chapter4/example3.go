package main

import "fmt"

type Inter interface {
	Ping()
	Pang()
}

type Anter interface {
	Inter
	String()
}

type St struct {
	Name string
}

func (St) Ping() {
	println("ping")
}

func (*St) Pang() {
	println("pang")
}

func (St) String() {
	println("string")
}

func main() {
	st := &St{"andes"}
	var i interface{} = st

	if o, ok := i.(Inter); ok {
		o.Ping() // Ping
		o.Pang() // Pang
	}

	if p, ok := i.(Anter); ok {
		// i 没有实现接口Anter，所以程序不会执行到这里
		p.String()
	}

	// 判断i绑定的实例是否就是具体类型St
	if s, ok := i.(*St); ok {
		fmt.Printf("%s\n", s.Name) // andes
	}
}
