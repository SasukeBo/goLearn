package main

import "fmt"

func main() {
	mp := make(map[int]string)
	mp[1] = "tom"
	mp[1] = "pony"
	mp[2] = "jaky"
	mp[3] = "andes"
	delete(mp, 3)

	fmt.Println(mp[1])
	fmt.Println(len(mp))

	for k, v := range mp {
		fmt.Println("key = ", k, " value = ", v)
	}
}
