package main

import "fmt"

// 变量的作用域
var x = 100 // 定义全局变量

// 定义一个函数
func f1() {
	// 函数中查找变量的顺序
	// 1.先在函数内部查找
	// 2.找不到在函数外面找，一直找到全局变量
	fmt.Println(x)
}

func main() {
	f1()
}
