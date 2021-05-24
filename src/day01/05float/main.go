package main

import "fmt"

// 浮点型
func main() {
	f1 := 1.23456 // 32位浮点数最大值
	fmt.Printf("%T \n", f1)
	// 默认go语言中的小数都是float64

	发 := float32(1.2345) // 显示声明为float32类型
	fmt.Printf("%T \n", 发)

	// 发 = f1 float32不能强制赋值给float64

}
