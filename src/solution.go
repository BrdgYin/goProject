package main

import "fmt"

var a []int
var n, k int

// 结构体中遇到的问题
func dfs(i int, sum int) bool {
	// 如果前n项都计算过了，则返回sum是否与K相等
	if i == n {
		return sum == k
	}
	// 不加上a[i]的情况
	if dfs(i+1, sum) {
		return true
	}

	// 加上a[i]的情况
	if dfs(i+1, sum+a[i]) {
		return true
	}

	// 无论是否加上a[i]都不能凑成k就返回false
	return false
}
func main() {
	a = []int{1, 2, 4, 7}
	n = 4
	k = 13
	if dfs(0, 0) {
		fmt.Println("yes")
	} else {
		fmt.Println("No")
	}
}
