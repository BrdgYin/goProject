package main

import "fmt"

// 结构体是值类型

type person struct {
	name, gender string
}

// Go和Java一样都是拷贝
// Java对于基本类型是值传递
// Java对于对象时拷贝对象的引用
func f(x *person) {
	x.gender = "女"
}
func main() {
	var p person
	p.name = "李四"
	p.gender = "男"
	fmt.Println(p)
	f(&p)
	fmt.Println(p)
}
