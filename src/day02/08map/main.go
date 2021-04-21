package main

import (
	"fmt"
)

// map

func main() {
	var m1 map[string]int // 需要初始化
	// 初始化，避免在程序运行期间动态扩容
	m1 = make(map[string]int, 10)
	m1["第一"] = 100
	m1["第二"] = 200
	m1["第3"] = 300
	m1["第4"] = 400
	m1["第5"] = 500
	// fmt.Println(m1)
	// 如果key不存在返回的是value对应类型的零值

	score, ok := m1["ji"] // 如果返回值是bool类型，通常用ok去接收
	if !ok {
		fmt.Println("没有数据")
	} else {
		fmt.Println("分数：", score)
	}
	delete(m1, "第一") // 如果删除的key不在什么都不干

	for k, v := range m1 { // 按照key, value的顺序遍历map
		fmt.Println("key: ", k, ", value: ", v)
	}

	for k := range m1 { // 按照key遍历map
		fmt.Println(k)
	}
	delete(m1, "kkkk")
}
