package main

import "fmt"

// switch 简化大量的判断

func main() {
	var n = 5
	switch n {
	case 1:
		fmt.Println("1")
		break
	case 2:
		fmt.Println("2")
		break
	case 3:
		fmt.Println("3")
		break
	case 4:
		fmt.Println("4")
	case 5:
		fmt.Println("5")
	default:
		fmt.Println("找不到")
	}

	switch n := 8; n {
	// 跟python中的写法一样
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")

	// 可以使用表达式
	case n % 2:
		fmt.Println("奇数")

	}
}
