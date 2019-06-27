package main

import "fmt"

func main() {
	mp := make(map[string]string, 10)
	mp["hello"] = "world"
	mp["你好"] = "世界"
	for key, value := range mp {
		fmt.Println("index is: ", key, " value is: ", value)
	}
}
