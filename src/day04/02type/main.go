package main

import "fmt"

// 自定义类型和类型别名

//type后面跟的是类型
type myInt int  // 自定义类型
type yInt = int // 类型别名

func main() {
	var c rune // int32--虽然是字符
	c = '中'
	fmt.Printf("%T\n", c)
}
