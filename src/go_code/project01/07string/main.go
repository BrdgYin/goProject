package main

import (
	"fmt"
	"strings"
)

// Go语言中字符串使用双引号包裹的
// Go语言中单引号包裹的是字符

func main() {
	// 字符串
	s := "Hello 南山"

	// 单独的字母、汉字、符号表示一个字符
	c1 := 'h'
	c2 := '1'
	c3 := '沙'
	c4 := "abcdefghijklmnopqrstuvwxyz"
	s2 := []string{"a", "b", "C"}

	// 1字节：1byte = 8bit
	fmt.Println(len(s))
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(strings.Index(c4, "c"))
	// 2
	fmt.Println(strings.Join(s2, ","))
	// a,b,c
}
