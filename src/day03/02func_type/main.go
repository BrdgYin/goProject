package main

import "fmt"

// 函数类型

func f1() {
	fmt.Println("hello")
}

func f2(i int) int {
	return 10
}

// 函数也可以作为参数传递
func f3(x func(int) int) {
	y := x(2)
	fmt.Println(y)
}

// 函数也可以作为返回值
func f5(x func()) func(int, int) int {
	// 执行函数x1
	x()
	// 定义的内部匿名函数
	ret := func(a, b int) int {
		return a + b
	}
	return ret
}

func main() {
	a := f1
	fmt.Printf("%T\n", a)
	// func()

	b := f2
	fmt.Printf("%T\n", b)
	// func(int) int

	f3(b)

	c := f5(f1)
	fmt.Println(c(1, 2))

}
