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
	_, ok := s.allStudent[id]
	if ok {
		fmt.Println("已存在当前用户")
		// 函数返回
		return
	} else {
		// 2.存在放入map
		s.allStudent[id] = student{
			id:   id,
			name: name,
		}
		fmt.Println("用户添加成功")
	}
}

// 删除学生
func (s studentMgr) deleteStudent() {
	// 获取输入
	fmt.Println("请输入要删除的id: ")
	var id int64
	fmt.Scanln(&id)
	// 直接查询
	_, ok := s.allStudent[id]
	if !ok {
		fmt.Printf("不存在该学生")
	} else {
		delete(s.allStudent, id)
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

	// 直接去map中取
	_, ok := s.allStudent[id]

	if !ok {
		fmt.Println("不存在这样的用户")
	} else {
		// 创建一个对象给它
		s.allStudent[id] = student{
			id:   id,
			name: name,
		}
		fmt.Println("修改成功")
	}
}
