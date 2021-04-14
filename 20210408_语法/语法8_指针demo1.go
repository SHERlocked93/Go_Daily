package main

import "fmt"

func main() {
	var a *int
	a = new(int) // 指针类型初始化使用 new 关键字
	fmt.Println(a, *a)

	var b map[string]int
	b = make(map[string]int, 8) // make 使用到 slice、map、channel 的初始化
	b["呵呵"] = 10
	fmt.Printf("%#v", b)
}
