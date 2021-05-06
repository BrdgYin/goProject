package main

import "fmt"

// 方法

// 标识符:变量名 函数名 类型名 方法名
// Go语言中如果标识符首字符是大写的，就表示对外可见(暴露的，公有的)
// 对于共有的注释格式为

// dog 这是一个狗的结构体
type dog struct {
	name string
}

// 构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

// 方法是作用于特定类型的函数
// d是调用者[接受者默认用首字母的小写]
// 指针也可以调用
func (d dog) wang() {
	fmt.Printf("%s, 汪汪 \n", d.name)
}

func main() {
	d1 := newDog("狗12") // 生成实例对象
	d1.wang()           // 调用方法
}
