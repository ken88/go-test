package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 读取文件
func main() {

	// 通过bufio 流方式读取
	bufioReadFile()

	//通过read 流方式读取
	// readFile()

	// 通过ioutil 方式读取
	// ioutilFile()
}

// 通过ioutil 方式读取
func ioutilFile() {
	byteStr, err := ioutil.ReadFile("/Users/yukai/Documents/go/go-test/aaa.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(byteStr))
}

// 通过bufio 流方式读取
func bufioReadFile() {
	file, err := os.Open("/Users/yukai/Documents/go/go-test/aaa.txt")
	defer file.Close() // 方法执行完毕 关闭文件
	if err != nil {
		fmt.Println(err)
		return
	}
	var fileStr string
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') // 表示一次读一行
		if err == io.EOF {
			fileStr += str
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fileStr += str
	}
	fmt.Println(fileStr)
}

//通过read 流方式读取
func readFile() {
	// 1.打开文件
	// 只读方式打开文件
	file, err := os.Open("/Users/yukai/Documents/go/go-test/aaa.txt")
	defer file.Close() // 方法执行完毕 关闭文件
	if err != nil {
		fmt.Println(err)
		return
	}
	// 2. 读取文件的内容
	var strSlice []byte
	var tempSlice = make([]byte, 128)
	for {
		n, err := file.Read(tempSlice) // 每次读取128字节
		if err == io.EOF {             // io.EOF 表示读取完毕
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println("读取失败")
		}
		// fmt.Printf("读取到了%v个字节\n", n)
		strSlice = append(strSlice, tempSlice[:n]...)
	}

	fmt.Println(string(strSlice))
}
