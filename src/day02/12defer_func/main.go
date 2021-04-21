package main

import "fmt"

// defer
// 第一步：返回值赋值
// [defer执行的时间]
// 第二步：真正的RET返回

func deferDemo() {
	fmt.Println("start")
	// defer把后面的语句延迟到函数即将返回的时候再执行

	// 执行顺序和栈一样
	defer fmt.Println("哈哈哈")
	defer fmt.Println("嘿嘿嘿")
	defer fmt.Println("呵呵呵")
	fmt.Println("end")
}

// 没有命名的返回就是x本身
func f1() int {
	x := 5
	defer func() {
		x++ // 修改的是x，不是返回值
	}()
	return x
}

// 返回值是x指向的值
func f2() (x int) {
	defer func() {
		x++
	}()
	// 返回值赋值 x = 5
	// 执行defer x ++
	// return   x = 6
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	// y = x
	// x ++
	// y = 5
	return x
}

func f4() (x int) {
	defer func(x int) {
		x++ // 改变的是函数的副本自己定义的x
	}(x) // 当作参数传递进来
	return 5 // 返回值 = x = 5
}

func f6() (x int) {
	defer func(x *int) {
		(*x)++
	}(&x)
	return 5 // 1.返回值 = x= 5 2. defer x = 6 3.ret x
}
func main() {
	deferDemo()
	// start
	// end
	// 呵呵呵
	// 嘿嘿嘿
	// 哈哈哈
	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 6
	fmt.Println(f3()) // 5
	fmt.Println(f4()) // 5
	fmt.Println(f6()) // 5
}
