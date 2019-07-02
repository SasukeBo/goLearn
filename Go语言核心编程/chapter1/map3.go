package main

import "fmt"

func main() {
	type User struct {
		name string
		age  int
	}

	ma := make(map[int]User)
	user := User{
		name: "sasuke",
		age:  24,
	}

	ma[1] = user
	// ma[1].age = 19 // ERROR, 不能通过map引用直接修改结构体成员
	user.age = 25
	ma[1] = user // 必须整体替换value

	fmt.Printf("user map is: %v\n", ma)
	fmt.Printf("user name is: %v\n", ma[1].name)
	fmt.Printf("user age is: %v\n", ma[1].age)
}
