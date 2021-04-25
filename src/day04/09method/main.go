package main

import "fmt"

// 给自定义类型加方法
// 不能给其他包中定义方法，只能给自己的包里的类型添加方法
type myInt int

func (i myInt) hello() {
	fmt.Println("给int方法添加方法")
}

func main() {
	m := new(myInt) // 使用new来实例化
	m.hello()
}
