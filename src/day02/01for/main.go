package main

import "fmt"

// 流程控制之前跳出for循环
func main() {
	// 当i = 5时跳出for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("for break;")
}
