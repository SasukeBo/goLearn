package main

func fa(base int) (func(int) int, func(int) int) {
	println(&base, base)

	add := func(i int) int {
		base += i
		println(&base, base)
		return base
	}

	sub := func(i int) int {
		base -= i
		println(&base, base)
		return base
	}

	return add, sub
}

func main() {
	f, g := fa(0)
	s, k := fa(0)
	println()

	println(f(1), g(2))
	println(f(1), g(2))
	println(s(1), k(2))
}
