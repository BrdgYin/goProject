package main

import "fmt"

// 构造函数

type person struct {
	name string
	age  int
}

// 构造函数
// 返回的是结构体还是结构体指针
// 当结构体占用的空间大时，尽量使用结构体指针，减少程序的内存开销
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

type dog struct {
	name string
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func main() {
	p1 := *newPerson("哈哈", 12)
	p2 := *newPerson("太卷了", 12)

	fmt.Println(p1)
	fmt.Println(p2)

	d1 := newDog("狗")
	fmt.Println(d1)
}
