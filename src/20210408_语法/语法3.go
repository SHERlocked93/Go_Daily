package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var scoreMap = make(map[string]int, 100)
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%02d", i)
		val := rand.Intn(100)
		scoreMap[key] = val
	}
	// 声明一个切片，保存所有key
	scoreStus := make([]string, 0, 100)
	for key := range scoreMap {
		scoreStus = append(scoreStus, key)
	}

	// 把切片中的key排序，并输出
	sort.Strings(scoreStus)
	for _, stu := range scoreStus {
		fmt.Println(stu, scoreMap[stu])
	}
}
