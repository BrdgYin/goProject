package main

import "fmt"

// 切片--简化版的动态数组
// 切片本质是引用类型是不能保存值的，值时放在底层内存中的数组中的

func main() {
	// 切片
	var s1 []int
	// 没有分配内存nil--所以为nil
	fmt.Println(s1 == nil)

	// 初始化方式1
	// s1 = []int{1, 2, 3}
	// fmt.Println(s1)

	// make初始化
	// 分配内存 长度 容量
	s2 := make([]bool, 2, 4)
	fmt.Println(s2)

	// 分配了内存所以不为nil
	s3 := make([]int, 0, 4)
	fmt.Println(s3 == nil)

	s4 := []int{1, 2, 3}
	s5 := s4
	fmt.Println(s4)
	s5[1] = 200
	fmt.Println(s4)
}
