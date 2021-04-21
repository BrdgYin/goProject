package main

import "fmt"

// 匿名函数
var f1 = func(x, y int) {
	fmt.Println(x + y)
}

func main() {
	f1(10, 20)
	// 函数内部无法声明带名字的函数

	f2 := func(x, y int) {
		fmt.Println(x + y)
	}

	f2(20, 30)

	// 如果只是调用一次的函数，还可以立即执行函数
	func() {
		fmt.Println("hello world")
	}() // 加括号立即执行
}
