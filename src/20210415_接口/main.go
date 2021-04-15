package main

import "fmt"

type sayer interface {
	say()
}

type cat struct{}
type dog struct{}
type person struct{ name string }

type xxx interface{}

func (c cat) say() {
	fmt.Println("miao~")
}
func (d dog) say() {
	fmt.Println("wang!")
}
func (p person) say() {
	fmt.Println("操？")
}

func da(arg sayer) {
	arg.say()
}
func main() {
	var a xxx
	a = cat{}
	fmt.Println(a)

	d1 := dog{}
	da(d1)

	p1 := person{name: "靓仔"}
	da(p1)
}
