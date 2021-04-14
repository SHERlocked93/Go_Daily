package main

import "fmt"

func main() {
	/* 单层循环 */
	//var cityArr = [...]string{"上海", "北京", "广州", "南京"}
	//for i := 0; i < len(cityArr); i++ {
	//	fmt.Println(cityArr[i])
	//}
	//for _, val := range cityArr {
	//	fmt.Println(val)
	//}

	/* for range */
	//cityArr2 := [...][2]string{
	//	{"上海", "北京"},
	//	{"广州", "深圳"},
	//	{"南京", "杭州"},
	//	{"苏州", "合肥"},
	//}
	//for _, val1 := range cityArr2 {
	//	fmt.Println(val1)
	//}
	numArr := [...]int{1, 2}
	//var sum int
	//for i := 0; i < len(numArr); i++ {
	//	sum += numArr[i]
	//}
	//fmt.Println(sum)
	numArr[0] = 5
	numsCopy := numArr
	numsCopy[0] = 6
	fmt.Printf("%v\t", numArr)
	fmt.Printf("%v\t", numsCopy)

	a := []int{1, 2, 3}
	b := a
	var c []int
	a[0] = -1
	copy(c, b)
	fmt.Println(a, b, c)
}
