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
func main() {
	a := f1
	fmt.Printf("%T\n", a)
	// func()

	b := f2
	fmt.Printf("%T\n", b)
	// func(int) int

	f3(b)

}
