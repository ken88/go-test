package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 写入文件
func main() {

	// 通过bufio 流方式写入
	// bufioWriteFile()

	//通过wrtie 流方式写入
	// writeFile()

	//  通过ioutil 方式写入
	ioutilWrite()
}

// 通过ioutil 方式写入 只能写入一条，每次都替换不能追加
func ioutilWrite() {
	str := "hello golang123"
	err := ioutil.WriteFile("/Users/yukai/Documents/go/go-test/test.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

}

// 通过bufio 流方式写入
func bufioWriteFile() {
	file, err := os.OpenFile("/Users/yukai/Documents/go/go-test/test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close() // 关闭文件
	if err != nil {
		fmt.Println(err)
		return
	}
	write := bufio.NewWriter(file)
	write.WriteString("bufio写入的数据\n")
	write.Flush()
}

// 通过wrtie
func writeFile() {
	/*
		os.O_WRONLY 只写
		os.O_CREATE 创建文件
		os.O_RDONLY 只读
		os.O_RDWR 读写
		os.O_TRUNC 清空
		s.O_APPEND 追加

	*/
	file, err := os.OpenFile("/Users/yukai/Documents/go/go-test/test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close() // 关闭文件
	if err != nil {
		fmt.Println(err)
		return
	}
	// 写入文件
	file.WriteString("写入的内容\n")
}
