package main

import "fmt"

func main() {
	a := 10
	b := &a
	fmt.Println(a, &a, b)
	fmt.Printf("%T %T", a, b) // int *int
}
