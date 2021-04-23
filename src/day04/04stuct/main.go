package main

// 结构体是值类型

type person struct {
	name, gender string
}

func main() {
	var p person
	p.name = "李四"
	p.gender = "女"
}
