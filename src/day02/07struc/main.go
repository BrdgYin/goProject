package main

import "fmt"

// 结构体

type Saiyan struct {
	Name  string
	Power int
}

// Go中传递参数到函数的方式: 镜像复制

func Super(s Saiyan) {
	s.Power += 10000
}

// 传递的是地址

func Super1(s *Saiyan) {
	s.Power += 10000
}

func main() {
	// 创建结构体
	goku := Saiyan{
		Name:  "Goku",
		Power: 9000, //注意这个逗号必须加
	}

	fmt.Println(goku)
	// {Goku 9000}
	Super1(&goku)
	fmt.Println(goku)
	// {Goku 19000}

	// 其它风格
	goku1 := Saiyan{}
	goku1.Name = "goku1"
	goku1.Power = 9000

	goku2 := Saiyan{"goku2", 2}
	fmt.Println(goku2)
	// {goku2 2}

	scores := make([]int, 5)
	scores = append(scores, 9332)
	fmt.Println(scores)
	// [0 0 0 0 0 9332]

	// 在前面插入9332
	scores1 := make([]int, 5)
	scores1 = append([]int{9332}, scores1...)
	fmt.Println(scores1)

}
