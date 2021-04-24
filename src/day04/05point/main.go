package main

import "fmt"

func main() {
	var a int
	a = 100
	b := &a

	fmt.Printf("type a:%T type b:%T\n", a, b)
	// type a:int type b:*int
	fmt.Printf("%p\n", &a)
	// 0xc000014098
	fmt.Printf("%v\n", b)
	// 0xc000014098
	fmt.Printf("%p\n", b)
	// 0xc000014098
	// a 和 b都是地址
	fmt.Printf("%p\n", &b)
	// 0xc000006028
	// &b是取b的地址， 地址的地址
}
