package main

import "fmt"

// 结构体占用一块连续的内存空间

type x struct {
	a int8
	b int8
	c int8
}

func main() {
	m := x{
		a: int8(10),
		b: int8(11),
		c: int8(12),
	}
	fmt.Printf("%d\n", &(m.a))
	fmt.Printf("%d\n", &(m.b))
	fmt.Printf("%d\n", &(m.c))
}
