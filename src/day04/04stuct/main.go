package main

import "fmt"

// 结构体是值类型

type person struct {
	name, gender string
}

// Go和Java一样都是拷贝
// Java对于基本类型是值传递
// Java对于对象时拷贝对象的引用
func f(x *person) {
	// go中由于不支持对指针操作，所以可以不用写作: (*x).gender
	x.gender = "女"
}
func main() {
	var p person
	p.name = "李四"
	p.gender = "男"
	fmt.Println(p)
	f(&p)
	fmt.Println(p)
	var p2 = new(person)
	p2.name = "王五"
	p2.gender = "女"
	fmt.Println(p2)
	// 对应类型的指针 &{王五 女}
	fmt.Println(*p2)
	// output:{王五 女}
	fmt.Printf("%T\n", p2)
	// 得到的是一个结构体指针
	fmt.Printf("%p\n", p2)

	// person需要提前定义--声明并同时初始化
	// key-value初始化
	var p3 = person{
		name:   "哈哈",
		gender: "难",
	}
	fmt.Println(p3)

	// 使用值列表的方式初始化，值的顺序要和结构体定义时字段顺序一致
	p4 := person{
		"努力",
		"有没有出路",
	}

	fmt.Println(p4)

}
