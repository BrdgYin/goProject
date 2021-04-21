package main

import "fmt"

// 函数
// 函数的定义
func sum(x int, y int) int {
	return x + y
}

// 没有返回值
func fun1(x int, y int) {
	fmt.Println(x, y)
}

// 命名返回值，就相当于在函数中声明了一个变量
func fun2(x int, y int) (ret int) {
	ret = x + y
	// 可以写return ret或者不用写
	return
}

// 多个返回值--要加括号
func fun3(x int, y int) (int, string) {
	return x + y, fmt.Sprintf("%d + %d = %d", x, y, x+y)
}

// 参数的类型简写，当参数中连续参数的类型一致时，我们可以将相同类型的参数类型省略
func fun4(x, y int) int {
	return x + y
}

// 可变参数
// 表示y可以传多个值[可变参数--类型推导],必须放到参数的最后
func fun5(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
}

func main() {
	// 在一个命名的的函数中严禁写其他的命名函数
	r := sum(10, 20)
	fmt.Println(r)
	fun1(10, 20)
	r, rS := fun3(10, 20)
	fmt.Println(rS)

	fun5("可变参数", 1, 2, 3, 4, 5, 6)
	fun5("可变参数")
}
