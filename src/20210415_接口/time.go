package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	t2 := now.Add(time.Hour)

	fmt.Println(t2)
	fmt.Printf("%#v", t2)

	//for tmp := range time.Tick(time.Second) {
	//	fmt.Println(tmp)
	//}

	fmt.Println(now.Format("2006-01-02 03:04:05 PM"))

	timeStr := "2001-01-05 10:11:59"
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		fmt.Println("err!!!", err)
		return
	}
	fmt.Println(t1.Format("2006/01==02 03--04.000 PM"))
}
