package main

/*
 	函数版学生管理系统:
	 	查看、新增、删除
*/
import (
	"fmt"
	"os"
)

// student 学生结构体
type student struct {
	id   int64
	name string
}

// 构造函数
func newStudent(id int64, name string) student {
	return student{
		id:   id,
		name: name,
	}
}

var (
	allStudent map[int64]*student
)

// showAllStudent 显示全部学生
func showAllStudent() {
	for k, v := range allStudent {
		fmt.Printf("学号: %d, id: %d, 姓名: %s\n", k, v.id, v.name)
	}
}

// 判断学生是否存在
func existStudent(id int64) bool {
	for k, _ := range allStudent {
		if k == id {
			return true
		}
	}
	return false
}

// addStudent 添加学生
func addStudent() {
	// 1.创建一个学生、
	var (
		id   int64
		name string
	)
	// 1.1获取用户输入
	fmt.Print("学生的id: ")
	fmt.Scanln(&id)
	fmt.Print("学生姓名: ")
	fmt.Scanln(&name)
	// 判断该学生是否存在
	if existStudent(id) {
		fmt.Println("该学生已存在！")
		return
	}
	// 1.2构造学生对象
	newStu := newStudent(id, name)
	// 2.追加到allStudent中
	allStudent[id] = &newStu
}

// deleteStudent 删除学生
func deleteStudent() {
	// 1.请用户输入要删除的学生的序号
	var (
		deleteId int64
	)
	fmt.Print("输入删除的id: ")
	fmt.Scanln(&deleteId)
	if !existStudent(deleteId) {
		fmt.Println("该学生不存在！")
		return
	}
	// 2.去allStudent这个map中删除对应的键值对
	delete(allStudent, deleteId)

}
func main() {
	// 初始化一个空间
	allStudent = make(map[int64]*student, 50)
	for { // 死循环
		// 1. 打印功能
		fmt.Println("welcome to sms!")
		fmt.Printf("1.查看所有学生 \n2.新增学生 \n3.删除学生 \n4.退出系统\n")
		fmt.Print("请输入: ")
		var input int
		fmt.Scanln(&input)
		fmt.Printf("当前所选择的功能：%d\n", input)

		// 2. 等待用户选择
		switch input {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("输入错误")
		}
		// 3. 执行相应的功能
	}
}
