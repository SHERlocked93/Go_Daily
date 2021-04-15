package main

import "fmt"

func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

type person struct {
	name, city string
	age        int8
}

func main() {
	var xiaoMin = new(person)
	(*xiaoMin).name = "小明"
	xiaoMin.city = "北京"
	xiaoMin.age = 19
	fmt.Printf("%#v", xiaoMin)

	p1 := newPerson("老王", "北京", 12)
	fmt.Printf("%#v", p1)
}
