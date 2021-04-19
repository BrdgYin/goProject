package main

import "fmt"

// for循环

func main() {
	// 基本格式
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// 变种1
	// var i = 5
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// 死循环
	// for {
	// 	fmt.Println(123)
	// }

	// for range
	// s := "Hello world哈哈"
	// for i, v := range s {
	// 	fmt.Printf("%d %c\n", i, v)
	// }

	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d x %d = %2d  ", i, j, i*j)
		}
		fmt.Printf("\n")

	}
}
