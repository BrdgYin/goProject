package main

import "fmt"

// base int: 做运算的基数
func calc(base int) (func(int) int, func(int) int) {
	// 封闭在函数内部
	// add 与 sub 都是用的base所以base是全局的

	// base = base + add_i
	add := func(i int) int {
		base += i
		return base
	}

	// base = base + add_i - sub_i
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	add, sub := calc(10)
	fmt.Println(add(10))
	fmt.Println(sub(5))
	// 20
	// 15
}
