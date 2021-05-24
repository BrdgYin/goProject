package main

import "fmt"

// 学生管理系统
// 具有得方法：
// 			1.保存一些数据--属性
//          2.拥有得功能--方法
type student struct {
	id   int64
	name string
}

// 学生构造函数
func (s student) Student(id int64, name string) student {
	return student{
		id:   id,
		name: name,
	}
}

// 管理对象
type studentMgr struct {
	allStudent map[int64]student
}

// 查看学生
func (s studentMgr) showStudent() {
	// 遍历map容器
	for _, v := range s.allStudent {
		fmt.Printf("学生编号:%d 学生名称:%s \n", v.id, v.name)
	}
}

// 增加学生
func (s studentMgr) addStudent() {
	// 获取用户输入
	var (
		id   int64
		name string
	)
	fmt.Println("请输入学号: ")
	fmt.Scanln(&id)
	fmt.Println("请输入姓名: ")
	fmt.Scanln(&name)
	// 0.判断是否已存在
	for k := range s.allStudent {
		if k == id {
			fmt.Println("学生已存在！")
			// return 不是系统调用，也不是库函数，而是一个关键字，
			// 表示调用堆栈的返回（过程活动记录），是函数的退出，而不是进程的退出。
			return // 直接返回--终止函数
		}
	}
	// 1.根据用户输入的内容创建新的学生
	var newStu = student{
		id:   id,
		name: name,
	}

	// 3.存在放入map
	s.allStudent[newStu.id] = newStu
}

// 删除学生
func (s studentMgr) deleteStudent() {
	// 获取输入
	fmt.Println("请输入要删除的id: ")
	var id int64
	fmt.Scanln(&id)

	var flag = true
	// 默认是key
	for k := range s.allStudent {
		if k == id {
			delete(s.allStudent, id)
			flag = false
		}
	}
	if flag {
		fmt.Printf("不存在该学生")
	} else {
		fmt.Println("删除成功")
	}
}

// 修改学生
func (s studentMgr) editStudent() {
	// 定义变量
	var (
		id   int64
		name string
	)

	// 获取输入
	fmt.Println("请输入要修改得id: ")
	fmt.Scanln(&id)
	fmt.Println("请输入name: ")
	fmt.Scanln(&name)
	var flag = true

	for k, v := range s.allStudent {
		if k == id {
			v.name = name
			flag = false
		}
	}
	if flag {
		fmt.Println("不存在这样的用户")
	} else {
		fmt.Println("修改成功")
	}
}
