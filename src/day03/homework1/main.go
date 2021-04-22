package main

import (
	"fmt"
)

// 全局变量
var (
	coins = 50
	users = []string{"Matthew", "Sarah", "Augstus", "Heidi", "Emilie", "Peter",
		"Giana", "Adriano", "Aaron", "Elizabeth"}
)

// 分金币
func dispathCoin(str []string, coin int) (map[string]int, int) {
	var distribution = make(map[string]int, len(str))
	for _, name := range str {
		temp := 0
		for _, s := range name {
			switch s {
			case 'e', 'E':
				temp++
			case 'i', 'I':
				temp += 2
			case 'o', 'O':
				temp += 3
			case 'u', 'U':
				temp += 4
			}
		}
		coin -= temp
		distribution[name] = temp
	}
	return distribution, coin
}

func main() {
	distributions, result := dispathCoin(users, coins)
	fmt.Println(distributions)
	fmt.Println(result)
}
