package main

import "fmt"

// 方法

// 标识符:变量名 函数名 类型名 方法名
// Go语言中如果标识符首字符是大写的，就表示对外可见(暴露的，公有的)
// 对于共有的注释格式为
// Person 这是一个人的类型
type person struct {
	name string
	age  int
}

// 构造函数
// 返回的是结构体还是结构体指针
// 当结构体占用的空间大时，尽量使用结构体指针，减少程序的内存开销
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

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
// d是调用者[接受者默认用首字母的小写]--必需是这个对象
// 指针也可以调用
func (d dog) wang() {
	fmt.Printf("%s, 汪汪 \n", d.name)
}

// 由于go是拷贝传参，所以无法修改age
func (p person) older() {
	p.age++
}

/*需要统一格式--前面用指针的情况下，后面也需要用指针*/
// 调用的时候相当于传递指针
func (p *person) olderP() {
	p.age++
}

// 吃饭
// 保持所有方法格式的统一性:大多情况下都使用指针类型
func (p *person) eat() {
	fmt.Println("马上去吃饭")
}

func main() {
	d1 := newDog("狗12") // 生成实例对象
	d1.wang()           // 调用方法

	p1 := newPerson("人", 12)
	p1.older()
	fmt.Println(p1.age) // 12
	p1.olderP()
	fmt.Println(p1.age) // 13
	p1.eat()
}
