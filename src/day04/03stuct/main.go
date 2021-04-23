package main

import "fmt"

// 结构体

// type定义类型
// struct表示结构体
type Person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	// 声明一个结构体的类型
	var a Person

	// 通过字段赋值
	a.name = "张三"
	a.age = 10
	a.gender = "男"
	a.hobby = []string{"打游戏", "听歌"}

	fmt.Printf("%v\n", a)
	// 访问变量的字段
	fmt.Println(a.name)

	// 匿名结构体:临时的应用场景
	var point struct {
		x int
		y int
	}
	point.x = 100
	point.y = 100
	fmt.Println(point)

}
