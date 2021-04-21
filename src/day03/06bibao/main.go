package main

// 闭包
import (
	"fmt"
	"strings"
)

// suffix: string是文件后缀的名称
func makeSuffixFunc(suffix string) func(string) string {
	// name: string是文件具体的名称
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
func main() {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")

	fmt.Println(jpgFunc("test"))
	fmt.Println(txtFunc("test"))
}
