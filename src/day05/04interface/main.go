package main

import "fmt"

// 接口定义
type car interface {
	// 方法名(参数1, 参数2, ...)(返回类型1, 返回类型2, ...)
	// 一个变量如果实现了接口中规定的所有方法，那么就实现了这个接口(Java定义一样)
	run() // 对run方法的上层抽象--调用时自动向下转型决定调用的具体方法
}

type BMW struct {
	brand string
}

type Porsche struct {
	brand string
}

// 类似于模板设计模式--可以自动决定调用的方法
func drive(c car) {
	c.run()
}

// 所有的run方法
func (b BMW) run() {
	fmt.Printf("%s speed: 50\n", b.brand)
}

func (p Porsche) run() {
	fmt.Printf("%s speed: 100\n", p.brand)
}

func main() {
	var b1 = BMW{
		brand: "宝马",
	}
	var p1 = Porsche{
		brand: "保时捷",
	}

	drive(b1)
	drive(p1)
}
