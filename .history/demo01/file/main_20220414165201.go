package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// 通过ioutil复制文件
	// ioutilCopyFile()

	// 创建目录
	// createCataloguee()

	// 删除文件
	// deletefile()

	// 重命名
	// chongmingming()

	// 遍历文件夹
	dirname := "/Users/yukai/Documents/go/go-gin_demo"
	listFiles(dirname, 0)

}

// 遍历文件
func listFiles(dirname string, level int) {
	// level 用来记录当前递归的层数，生成带有层级感的空格
	s := "|--"
	for i := 0; i < level; i++ {
		s = "| " + s
	}
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range fileInfos {
		filename := dirname + "/" + fi.Name()
		fmt.Printf("%s%s\n", s, filename)
		if fi.IsDir() {
			// 递归调用方法
			listFiles(filename, level+1)
		}
	}
}

// 重命名
func chongmingming() {
	err := os.Rename("./a11.txt", "./aaa.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("重命名成功")
}

// 删除文件
func deletefile() {
	// 删除文件
	// err := os.Remove("./test.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("删除成功")

	// 删除/aaa/bbb 目录下的所有目录
	err := os.RemoveAll("./aaa/bbb")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("删除成功")
}

// 创建目录
func createCataloguee() {
	// err := os.Mkdir("./dir", 0755)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	err := os.MkdirAll("./dir/dir1/dir2", 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 通过ioutil复制文件
func ioutilCopyFile() {
	byteStr, err1 := ioutil.ReadFile("/Users/yukai/Documents/go/go-test/aaa.txt")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	err2 := ioutil.WriteFile("/Users/yukai/Documents/go/go-test/test1.txt", byteStr, 0666)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println("复制成功")
}
