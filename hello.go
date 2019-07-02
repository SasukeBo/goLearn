package main

import "fmt"

func main() {
	str := "Lbh penpxrq gur pbqr!"
	strclice := []byte(str)

	for i := 0; i < len(strclice); i++ {
		ch := strclice[i]
		switch {
		case (ch > 97) && (ch < 122):
			strclice[i] = (ch-97+13)%26 + 97

		case (ch > 65) && (ch < 90):
			strclice[i] = (ch-65+13)%26 + 65

		default:
			strclice[i] = ch
		}
	}

	fmt.Println(string(strclice))
}
