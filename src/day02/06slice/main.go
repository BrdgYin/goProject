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

	var s6 []int
	// 自动初始化切片
	s6 = append(s6, 1)
	
	// var s7 []int // 为nuil --[]
	// var s7 = make([]int, 0, 3) // 由于此时长度为0 --[]
							   // 然而copy函数不会扩容所以是[]
	var s7 = make([]int, 3)    // 此时开辟了3个空间 -- [1 0 0]
	copy(s7, s6)
	fmt.Println(s7)

	// 指针
	// Go里面的指针只能读取不能修改，不能修改指针变量指向的地址
	addr := "重庆邮电大学信科大厦"
	addrP := &addr
	fmt.Println(addrP)
	fmt.Printf("%T\n", addrP)
	// 0xc000088230
	// *string
	addrV := *addrP
	fmt.Println()
	fmt.Println(addrV)
}
