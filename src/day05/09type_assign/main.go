package main

import (
	"fmt"
)

// 类型断言

func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	// 接口.(试图转换的类型)
	// ok == true-->str就是对应的值
	// ok == false->str就是个0值
	str, ok := a.(string)
	if !ok {
		fmt.Println("Error")
		return
	}
	fmt.Println("传进来的是: ", str)
}

func assign2(a interface{}) {
	fmt.Printf("%T\n", a) // 类似于这种方法
	switch t := a.(type) {
	case string:
		fmt.Println("字符串: ", t)
	case int:
		fmt.Println("int: ", t)
	case bool:
		fmt.Println("bool: ", t)
	}
}

func main() {
	assign(111)
	assign2("1111")
}
