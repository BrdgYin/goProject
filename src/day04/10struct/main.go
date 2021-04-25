package main

import "fmt"

// 结构体中的问题
type myInt int

func (m myInt) hello() {
	fmt.Println("自定义的int")
}

func main() {
	// 这里把myInt当做一个结构体
	m := myInt(100)
	m.hello()
}
