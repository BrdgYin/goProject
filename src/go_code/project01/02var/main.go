package main

import "fmt"

// 声明变量--需要每次都写一个var
var name string
var age int
var isOk bool

// 批量声明--此时为对应的空值
var (
	name1 string // 对应""
	age1  int    // 0
	isOk1 bool   // false
)

func main() {
	var studentName string
	name = "理想"
	age = 16
	isOk = true
	studentName = "学生姓名"

	// 声明变量同时赋值
	var studentAge int = 11

	fmt.Printf("name:%s \n", name)
	fmt.Printf("age:%d  \n", age)
	fmt.Println("studentName: ", studentName)
	fmt.Println("studentAge: ", studentAge)
	// 类型推导(根据值判断该变量是什么类型)
	var s2 = "20"
	fmt.Println(s2)

	// 短变量声明--只能声明新的变量
	// 只能在函数内部，可以使用更简洁的:=方式初始化变量
	s3 := "哈哈哈"
	fmt.Println(s3)

	// 匿名变量"_":python--占位

}
