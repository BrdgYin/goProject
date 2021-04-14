package main

// 表示该main.go文件所在的包是main，每个文件都必须属于一个包

import (
	"fmt"
)

// 引入一个包，引入一个包之后就可以使用包中的函数
// 类似于include/import

// 函数外只能放置标识符(变量/函数/类型)声明

func main() { // 一个包只有一个main
	fmt.Println("hello world")
}
