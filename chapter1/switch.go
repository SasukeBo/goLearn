package main

import "fmt"

func main() {
	switch i := "y"; i { // switch后面可以带上一个初始化语句
	case "y", "Y": // 多个case值使用逗号分隔
		fmt.Println("yes")
		fallthrough

	case "n", "N":
		fmt.Println("no")
	}

	score := 85
	grade := ' '

	if score >= 90 {
		grade = 'A'
	} else if score >= 80 {
		grade = 'B'
	} else if score >= 70 {
		grade = 'C'
	} else if score >= 60 {
		grade = 'D'
	} else {
		grade = 'F'
	}

	// 上面的if else 可以改写为下面的switch语句
	switch {
	case score >= 90:
		grade = 'A'
	case score >= 80:
		grade = 'B'
	case score >= 70:
		grade = 'C'
	case score >= 60:
		grade = 'D'
	default:
		grade = 'F'
	}

	fmt.Printf("grade = %c\n", grade)
}
