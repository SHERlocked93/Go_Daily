package main

import "fmt"

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 200
}

func main() {
	a, b := 1, 2
	modify1(a)  // 1
	modify2(&b) // 200
	fmt.Println(a, b)
}
