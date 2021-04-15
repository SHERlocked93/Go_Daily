package calc

import "fmt"

var Name = "帅哥"

func Add(x, y int) int {
	return x + y
}

func init() {
	fmt.Println("calc init")
}
