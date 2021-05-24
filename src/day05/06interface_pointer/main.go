package main

/*
 * 使用值接收者实现接口于使用指针接收者实现接口得区别:
 * 使用值接收者实现接口: 结构体类型和结构体指针类型得变量都能存
 * 指针接收者实现接口只能存结构体指针类型得变量
 */
import "fmt"

// 使用值接受者和使用指针接受者的区别
type animal interface {
	move()
	eat(something string)
}

type cat struct {
	Name string
	Feet int8
}

// 使用值接受者实现了接口得所有方法
// func (c cat) move() {
// 	fmt.Println("跑跑跳跳")
// }

// func (c cat) eat(s string) {
// 	fmt.Printf("猫吃老鼠 \n")
// }

// 使用指针接收者实现了所有得接口方法
func (c *cat) move() {
	fmt.Println("跑跑跳跳")
}

func (c *cat) eat(s string) {
	fmt.Printf("猫吃老鼠 \n")
}

func main() {
	var a1 animal

	c1 := cat{"tom", 4}
	c2 := &cat{"jom", 3} // 获取指针类型

	fmt.Printf("%T\n", c2)
	fmt.Printf("%T\n", c1)

	// *main.cat
	// main.cat
	// 值接收者得使用方法
	// a1 = c1

	// 使用指针接收者, 需要获取结构体对象得地址
	a1 = &c1
	fmt.Printf("%T\n", a1)
	fmt.Printf("%v\n", a1)

	a1.eat("呵呵")
	a1.move()

	a1 = c2
	fmt.Printf("%T\n", a1)
	fmt.Printf("%v\n", a1)
	// main.cat
	// {tom 4}
	// *main.cat
	// &{jom 3}

	a1.eat("哈哈")
	a1.move()

}
