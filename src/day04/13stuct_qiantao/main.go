package main

import "fmt"

// 嵌套结构体

type person struct {
	name string
	age  int
	addr address // 结构体嵌套
}

type company struct {
	name    string
	address // 匿名结构体嵌套
}

type address struct {
	province string
	city     string
}

func main() {
	// 未初始化的位置未空
	p1 := person{
		name: "大大",
		age:  100,
		addr: address{
			province: "四川",
			city:     "成都",
		}, // 这个地方必须要加","
	}

	c1 := company{
		name: "重庆邮电大学",
		address: address{
			province: "重庆市",
			city:     "重庆市南岸区",
		},
	}

	fmt.Println(p1)
	// 拿到特定的属性
	fmt.Println(p1.name, p1.addr.province)

	fmt.Println(c1)
	// 那特定的属性
	fmt.Println(c1.name, c1.address.city)
	// 由于名称唯一也可以直接拿取
	// 1.现在自己的结构体内找 2.在自己嵌套的结构体找 3.如果重名会报错--冲突
	fmt.Println(c1.city)
}
