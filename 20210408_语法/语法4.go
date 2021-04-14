package main

import (
	"fmt"
	"strings"
)

func main() {
	// 遍历一个句子中每个单词出现的次数
	str := "how do you do"
	strMap := make(map[string]int, 10)
	for _, word := range strings.Split(str, " ") {
		_, ok := strMap[word]
		if ok {
			strMap[word] ++
		} else {
			strMap[word] = 1
		}
	}
	fmt.Printf("%#v", strMap)
}
