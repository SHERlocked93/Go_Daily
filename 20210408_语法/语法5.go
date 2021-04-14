package main

import "fmt"

func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

func main() {
	x, y := calc(100, 200)
	fmt.Println(x, y)
}
