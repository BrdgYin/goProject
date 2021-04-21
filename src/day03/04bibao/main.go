package main

// 类型不匹配的函数传进来
import "fmt"

// 自己的代码逻辑
func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

// 别人的代码逻辑
func f2(x, y int) {
	fmt.Println("f2()")
	fmt.Println(x + y)
}

// 闭包
// 为了让f1可以在不改动的情况下调用f2
// 定义一个函数对f2进行包装
func f3(f2 func(int, int), m, n int) func() {
	fmt.Println("闭包")
	// 匿名函数--包括了要执行的f2
	// 相当于用包装了的tmp来执行想要执行的内容
	// 此时tmp的函数类型与f3的返回值类型一样--刚好是f1的参数类型
	tmp := func() {
		// 调用f2
		f2(m, n)
	}
	return tmp
}
func main() {
	f1(f3(f2, 100, 200))
	// 把原来需要传递两个int类型的参数包装成一个不需要传参的函数
}
