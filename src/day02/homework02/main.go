package main

import "fmt"

// 需要考虑是不是汉字
func F(str string) {
	r := make([]rune, 0, len(str))

	for _, c := range str {
		r = append(r, c)
	}
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			fmt.Println("false")
			return
		}
	}
	fmt.Println("true")
}
func main() {
	F("上海自来水来自海上")

}
