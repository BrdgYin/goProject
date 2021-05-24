package main

import "fmt"

// 接口实现
type animal interface {
	move()
	eat(something string)
}

type cat struct {
	Name string
	Feet int8
}

type chicken struct {
	Feet int8
}

func (c cat) move() {
	fmt.Println("跑跑跳跳")
}

func (c cat) eat(s string) {
	fmt.Printf("猫吃老鼠 \n")
}

func (c chicken) move() {
	fmt.Println("一摇一摆")
}

func (c chicken) eat(s string) {
	fmt.Printf("小鸡吃虫 \n")
}

func main() {
	var c = cat{
		Name: "Tom 猫",
		Feet: 4,
	}

	var ch = chicken{
		Feet: 2,
	}

	var a1 animal // 接口类型的变量
	a1 = c        // 将猫的对象赋值给a1
	fmt.Println(a1)
	a1.eat("哈哈")
	a1.move()
	// 但是不能调用属性--跟之前的jdk定义一样

	a1 = ch
	fmt.Println(a1)
	a1.eat("哈哈")
	a1.move()
	// 但是不能调用属性--跟之前的jdk定义一样

	/* 给不同的对象执行不同的方法
	{Tom 猫 4}
	猫吃老鼠
	跑跑跳跳
	{2}
	小鸡吃虫
	一摇一摆
	*/

}
