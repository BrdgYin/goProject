package main

import "fmt"

// 同一个结构体可以实现多个接口
// 接口就像一个标准和规则

type animal interface {
	// 接口嵌套
	mover
	eater
}
type mover interface {
	move()
}

type eater interface {
	eat()
}

// 同时实现了两个接口
type cat struct {
	name string
	feet int8
}

func (c *cat) move() {
	fmt.Println("蹦蹦跳跳")
}

func (c *cat) eat() {
	fmt.Println("eating ....")
}

func main() {

}
