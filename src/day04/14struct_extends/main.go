package main

import "fmt"

// 结构体模拟实现其他语言中的"继承"

type animal struct {
	name string
}

// 给animal实现行动的方法
func (a animal) move() {
	// 可以调用a中的方法
	fmt.Printf("%s 会动\n", a.name)
}

// 实现狗类
type dog struct {
	feet   uint
	animal // dog将所有的animal的方法包括进来了
}

// 给dog实现方法--叫
func (d dog) speak() {
	fmt.Printf("%s 汪汪\n", d.name)
}

func main() {
	d1 := dog{
		feet: 4,
		animal: animal{
			name: "小黑",
		},
	}

	d1.speak()
	// animal的方法dog也可以调用
	d1.move()
}
