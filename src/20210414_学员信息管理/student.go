package main

import "fmt"

type student struct {
	id    int // 唯一学号
	name  string
	class string
}

// 构造函数
func newStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

// 学员管理类型
type studentMar struct {
	allStudents []*student
}

// 学员管理 构造函数
func newStudentMar() *studentMar {
	return &studentMar{
		allStudents: make([]*student, 0, 100),
	}
}

// 添加学员
func (s *studentMar) addStudent(newStu *student) {
	s.allStudents = append(s.allStudents, newStu)
}

// 编辑学员
func (s *studentMar) modifyStudent(newStu *student) {
	for idx, stu := range s.allStudents {
		if stu.id == newStu.id {
			s.allStudents[idx] = newStu
			return
		}
	}
	fmt.Println("未能找到学生 ", newStu.name, "学号 ", newStu.id)
}

// 展示学员
func (s *studentMar) showStudent() {
	for _, v := range s.allStudents {
		fmt.Printf("学号:%d\t 姓名:%s\t 班级:%s\n", v.id, v.name, v.class)
	}
}
