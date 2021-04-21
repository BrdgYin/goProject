package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("10", a, b))
	b = 1
}

// 1.calc("1", a, calc("10", a, b))
// 2.calc("10", a, b) // 10 1 2 3
// 3.a = 0
// 4.calc("1", a, calc("10", a, b))
// 5.calc("10", 0, 2) // 10 0 2 2
// 6.先进后出 calc("2", a, 2) // 2 0 2 2
// 7.calc("1", a, 3) // 1 1 3 4
