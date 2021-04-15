package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readFromFile(path string) {
	fileObj, err := os.Open(path)
	if err != nil {
		fmt.Println("err!", err)
		return
	}

	defer fileObj.Close()
	var tmp = make([]byte, 128)
	n, err := fileObj.Read(tmp)
	if err != nil {
		fmt.Println("read err !", err)
		return
	}
	fmt.Printf("read byte: %d", n)
	fmt.Println(string(tmp))
}

func readBigFile(path string) {
	fileObj, err := os.Open(path)
	if err != nil {
		fmt.Println("err!", err)
		return
	}

	defer fileObj.Close()

	for {
		var tmp = make([]byte, 128)
		n, err := fileObj.Read(tmp)
		if err == io.EOF { // end of file
			fmt.Println(string(tmp))
			return
		}
		if err != nil {
			fmt.Println("read err !", err)
			return
		}
		fmt.Printf("read byte: %d \n", n)
		fmt.Println(string(tmp[:n]))
	}
}

func readByBufio(path string) {
	fileObj, err := os.Open(path)
	if err != nil {
		fmt.Println("err!", err)
		return
	}

	defer fileObj.Close()

	reader := bufio.NewReader(fileObj)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("reader err %v \n", err)
		return
	}
	fmt.Println(line)
}

func readBigByBufio(path string) {
	fileObj, err := os.Open(path)
	if err != nil {
		fmt.Println("err!", err)
		return
	}

	defer fileObj.Close()
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Print(line)
			return
		}
		if err != nil {
			fmt.Printf("reader err %v \n", err)
			return
		}
		fmt.Print(line)
	}
}

func writeString() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer fileObj.Close()
	str := "沙河小王子"
	fileObj.Write([]byte(str))      // []byte
	fileObj.WriteString("hello 沙河") // string
}

func main() {
	//readFromFile("src/20210416_文件/xx.txt")
	//readBigFile("src/20210416_文件/xx.txt")
	//readByBufio("src/20210416_文件/xx.txt")
	readBigByBufio("src/20210416_文件/xx.txt")
}
