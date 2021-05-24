package main

import "fmt"

// 整型

func main() {
	var i1 = 101
	fmt.Printf("%d \n", i1)
	fmt.Printf("%b \n", i1)

	// 八进制数 0开头
	i2 := 076
	fmt.Printf("%o \n", i2)
	fmt.Printf("%d \n", i2)

	// 十六进制数
	i3 := 0x1234567890
	fmt.Printf("%x\n", i3)

	// 查看变量的类型
	i4 := int16(0)
	fmt.Printf("%T \n", i4)

	/**
	 * 1100101
	 * 76
	 * 62
	 * 1234567890
	 * int16
	 */
}
