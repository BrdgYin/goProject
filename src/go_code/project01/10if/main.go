package main

import "fmt"

// 条件判断 https://www.bilibili.com/video/BV16E411H7og?p=16
func main() {
	// age := 19
	// if age > 18 {
	// 	fmt.Println("成年了")
	// } else {
	// 	fmt.Println("上课了")
	// }

	// // 多个条件判断
	// if age > 35 {
	// 	fmt.Println("养老了")
	// } else if age > 18 {
	// 	fmt.Println("没得了")
	// } else {
	// 	fmt.Println("上课了")
	// }

	// 特殊写法
	// age变量此时只在if条件判断语句中生效--[局部变量]
	if age := 19; age > 18 {
		fmt.Println("18")
	} else {
		fmt.Println("17")
	}
}
