package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	ID   int
	Name string
}

// student构造函数
func newStudent(ID int, Name string) student {
	return student{
		ID:   ID,
		Name: Name,
	}
}

type class struct {
	Title    string    `json:"title"`
	Students []student `json:"student_list"`
}

func main() {
	classA := class{ // 创建班级
		Title:    "火箭101",
		Students: make([]student, 0, 20),
	}

	for i := 0; i < 10; i++ {
		stu := newStudent(i, fmt.Sprintf("stu%02d", i))
		classA.Students = append(classA.Students, stu)
	}
	fmt.Printf("%#v\n", classA)

	// JSON 序列化
	strJson, err := json.Marshal(classA)
	if err != nil {
		fmt.Println("json marshal err:", err)
		return
	}
	fmt.Printf("%T\n %s\n\n", strJson, strJson)

	//	JSON 反序列化
	var dataJson class
	err2 := json.Unmarshal(strJson, &dataJson)
	if err2 != nil {
		fmt.Println("json unmarshal err:", err)
		return
	}
	fmt.Printf("%T\n %s\n", dataJson, dataJson)
}
