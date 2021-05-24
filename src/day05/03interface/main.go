package main

import "fmt"

// 接口实例--一种特殊的类型
// 不关心变量的具体类型，只关心调用的方法
type speaker interface {
	// 方法签名--可以有多个
	speak() // 类比java中的接口--多态
}

type cat struct{}

type dog struct{}

type person struct{}

func (c cat) speak() {
	fmt.Println("miao miao")
}

func (p person) speak() {
	fmt.Println("aaaaa")
}

func (d dog) speak() {
	fmt.Println("wang wang")
}

func fit(x speaker) {
	// 接受一个参数，传进来什么就执行什么
	x.speak()
}

func main() {
	var c person
	fit(c)
}
