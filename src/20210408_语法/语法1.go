package main

import "fmt"

func main() {
	var (
		age = 12
		a   = "stringA"
		b   string
		c   int
	)
	d := "ddd"

	fmt.Print(age, a, b, c, d)
	e := age
	fmt.Print(e, age)

	for i := 0; i < 10; i++ {
		//fmt.Println(i)
	}
	fmt.Println(len(d))
}
