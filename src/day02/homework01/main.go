package main

import (
	"fmt"
	"strings"
)

func findAlpha(str string) (res map[string]int) {
	var chas []byte = []byte(str)
	res = make(map[string]int, 10)
	pre := 0
	for i := 0; i < len(chas); i++ {

		// 统计带空格的
		if chas[i] == ' ' || i == len(chas)-1 {
			var ii int
			if i == len(chas)-1 {
				ii = i + 1
			} else {
				ii = i
			}
			val, ok := res[string(chas[pre:ii])]
			if ok == false {
				res[string(chas[pre:ii])] = 1
			} else {
				res[string(chas[pre:ii])] = val + 1
			}
			pre = i + 1 // 跳过空格
		}
	}
	return res
}

func findAlpha2(str string) map[string]int {
	res := make(map[string]int, 10) // 初始化并开辟空间
	var arr []string
	arr = strings.Fields(str)
	for _, val := range arr {
		if val != " " {
			i, ok := res[val]
			if ok == false {
				res[val] = 1
			} else {
				res[val] = i + 1
			}
		}
	}
	return res
}
func main() {
	fmt.Println(findAlpha2("how do you do   "))
}
