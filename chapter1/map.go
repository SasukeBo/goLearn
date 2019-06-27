package main

import "fmt"

func main() {
	mp1 := make(map[string]string)
	mp2 := make(map[string]int, 10)
	mp1["firstname"] = "汪"
	mp1["lastname"] = "波"
	mp2["age"] = 25

	fmt.Printf("姓名：%v", mp1["firstname"])
	fmt.Println(mp1["lastname"])

	fmt.Printf("年龄：%v\n", mp2["age"])
}
