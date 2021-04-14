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
	address
}

type address struct {
	province string
	cun      string
}

func (p person) Dream() {
	fmt.Printf("%s的梦想是天天收租", p.name)
}

func main() {
	p1 := newPerson("小明", "北京", 18)
	p1.Dream()
}
