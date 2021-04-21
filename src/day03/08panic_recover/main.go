package main

import "fmt"

// panic 和 recover
func funcA() {
	fmt.Println("a")
}
func funcB() {
	// 刚刚打开数据库连接
	// 设置一个立即执行的匿名函数
	defer func() {
		err := recover() // 尝试恢复现场
		fmt.Println(err)
		fmt.Println("释放掉打开的资源")
	}() // 有()表示立即执行
	panic("error:....") // 程序崩溃退出
	fmt.Println("b")
}
func funcC() {
	fmt.Println("c")
}

func main() {
	funcA()
	funcB()
	funcC()
	// a
	// 释放掉打开的资源
}
