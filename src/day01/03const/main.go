package main

import (
	"fmt"
)

// 常量

const PI = 3.1415926

// 批量声明常量
const (
	STATUSOK  = 200
	STATUSNTF = 404
)

const (
	N1 = 100
	N2 = 200
	N3 // 没有声明的下面默认和上面一样
	N4
)

// iota--类似于枚举
const (
	a1 = iota // 初次出现被初始化为  0
	a2 = iota // 同一个const，新增【一行】【常量声明】自动加一 1
	_         // 2
	a3        // 默认为iota 3
)

// 多个常量在【一行】
const (
	d1, d2 = iota + 1, iota + 2 // d1: 1 d2: 2 iota: 0
	d3, d4 = iota + 1, iota + 2 // d3: 2 d4: 3 iota: 1
)

// 定义数量级
const (
	_  = iota             // 初始为0
	KB = 1 << (10 * iota) // 1024 1 左移10位
	MB = 1 << (10 * iota) // 1024 * 1024 20位
	GB = 1 << (10 * iota) // 30位
	TB = 1 << (10 * iota) // 40位
	PB = 1 << (10 * iota) // 50位
)

func main() {
	// fmt.Println(N1, N2, N3, N4)
	/**
	 * 100 200 200 200
	 */

	// fmt.Println(1)
	// fmt.Println(2)
	// fmt.Println(3)
	/*
		0
		1
		2
	*/

	// fmt.Println(d1)
	// fmt.Println(d2)
	// fmt.Println(d3)
	// fmt.Println(d4)
	/**
	 * 1
	 * 2
	 * 2
	 * 3
	 */

	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(TB)
	/**
	* 1024
	* 1048576
	* 1099511627776
	 */
}
