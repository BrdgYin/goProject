package main

import "fmt"

// 结构体的匿名字段结构体嵌套
type person struct {
	string
	int
	// string 结构体内的属性必须时唯一的
}

func main() {
	p1 := person{
		"哈哈",
		111,
	}
	// 通过类型访问属性
	fmt.Println(p1.string)
	fmt.Println(p1.int)
}
