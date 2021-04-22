package main

import "fmt"

// 递归
func taijie(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return taijie(n-1) + taijie(n-2)
}
func main() {
	fmt.Println(taijie(3))
}
