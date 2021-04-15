package main

import "fmt"

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动\n", a.name)
}

type Dog struct {
	Feet int8
	*Animal
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪\n", d.name)
}

func main() {
	wangcai := &Dog{
		Feet:   4,
		Animal: &Animal{name: "旺财"},
	}
	wangcai.wang()
	wangcai.move()
	fmt.Printf("%#v %T", wangcai, wangcai)
}
