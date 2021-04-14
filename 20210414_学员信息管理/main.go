package main

import (
	"fmt"
	"os"
)

func showMenu() {
	fmt.Println("欢迎，选择你的操作 \n1. 添加学员\n2. 编辑学员\n3. 展示所有学员信息\n4. 退出")
}

// 获取用户输入信息
func getInput() *student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Println("输入 ID：")
	fmt.Scanf("%d\n", &id)
	fmt.Println("输入 Name：")
	fmt.Scanf("%s\n", &name)
	fmt.Println("输入 Class：")
	fmt.Scanf("%s\n", &class)
	return newStudent(id, name, class)
}

// 学员信息的 crud
func main() {
	sm := newStudentMar()
	for {
		showMenu()
		var opType int
		fmt.Scanf("%d\n", &opType)

		switch opType {
		case 1:
			newStu := getInput()
			sm.addStudent(newStu)
		case 2:
			newStu := getInput()
			sm.modifyStudent(newStu)
		case 3:
			sm.showStudent()
		case 4:
			os.Exit(0)
		}
	}
}
