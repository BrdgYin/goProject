package main

import "fmt"

// map和slice组合
func main() {
	// 元素类型为map的切片
	var s1 = make([]map[int]string, 2, 10)

	// 没有对内部的map初始话
	// s1[0][100] = "a"

	// 对map初始化
	s1[0] = make(map[int]string, 1)
	s1[0][1] = "A"
	fmt.Println(s1)

	// 值为切片的map
	var m1 = make(map[string][]int, 10)

	m1["1"] = []int{1, 2, 3}
	fmt.Print(m1)

	// [map[1:A] map[]]
	// map[1:[1 2 3]]
}
