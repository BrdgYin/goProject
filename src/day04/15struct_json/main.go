package main

import (
	"encoding/json"
	"fmt"
)

// 结构体于json

// 1.序列化: 把go语言中的结构体变量-->json格式的字符串
// 2.反序列化: json格式的字符串 --> Go语言中能够识别的结构体变量

type person struct {
	Name string `json:"name" db:"name" ini:"name"` // 可见性 大写才可见, + tag可以在json解析得时候指定命名
	Age  int    `json:"age"`                       // 可见性 大写才可见
}

func main() {
	p1 := person{
		Name: "哈哈",
		Age:  19,
	}
	// 序列化
	b, err := json.Marshal(p1)

	if err != nil {
		fmt.Printf("err: %v", err)
	}

	fmt.Printf("%v\n", string(b))

	// 反序列化
	str := `{"name": "张三", "age":20}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) // 传指针市方便函数去修改p2得值
	fmt.Printf("%#v\n", p2)
}
