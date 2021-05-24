package main

import "fmt"

// 空接口
// 空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。
// 空接口类型的变量可以存储任意类型的变量

func main() {
	// 声明并开辟空间
	var m1 = make(map[string]interface{}, 16)
	m1["name"] = "哈哈"
	m1["age"] = 100
	m1["merried"] = true

	// 数组类型
	// [...] 表示自动计算大小
	m1["hobby"] = [...]string{"唱", "跳", "rap"}

	fmt.Println(m1)

}
