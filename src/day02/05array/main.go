package main

import "fmt"

// 数组
// 存放元素的容器,必须指定存放元素的类型和容量
func main() {
	// 声明数组--类似java
	var a1 [3]int
	// 自动推断长度
	var a2 = [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	var s1 [3]string = [3]string{"成都", "重庆", "昆明"}
	a1[0] = 1
	a1[1] = 2
	a1[2] = 3

	// 多维数组
	var a11 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println("a2:", len(a2))

	// 遍历数组 类似于python enum
	for _, i := range s1 {
		fmt.Println(i)
	}
	for j := 0; j < len(a2); j++ {
		fmt.Println(a2[j])
	}
	for _, i := range a11 {
		for _, j := range i {
			fmt.Print(j, ",")
		}
		fmt.Println()
	}
}
