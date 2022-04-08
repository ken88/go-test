package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1.字符串转切片
	// str1 := "123-456-789"
	// res := strings.Split(str1, "-") // [123 456 789]
	// fmt.Println(res)

	// 2.切片转字符串拼接
	// str1 := "123-456-789"
	// res1 := strings.Split(str1, "-")
	// str := strings.Join(res1, "**")
	// fmt.Println(str) // 123**456**789 字符串转切片

	// arr := []string{"php", "java", "golang"}
	// str2 := strings.Join(arr, "-")
	// fmt.Println(str2) // php-java-golang

	// 3. strings.Contains 判断一个字符串中是否包含另一个字符串
	// str3 := "this is golang"
	// str4 := "golang"
	// flag := strings.Contains(str3, str4)
	// fmt.Println(flag)

	// 4.strings.HasPrefix 一个字串是否在另一个字符串前缀
	// str3 := "你好is123 is golang"
	// str4 := "你好"
	// flag := strings.HasPrefix(str3, str4)
	// fmt.Println(flag)

	// 5.strings.HasSuffix 一个字串是否在另一个字符串后缀 返回true  false
	// str3 := "你好is123 is golang你好123"
	// str4 := "你好123"
	// flag := strings.HasSuffix(str3, str4)
	// fmt.Println(flag)

	// 6. strings.Index  strings.LastIndex 下标从0开始 查找到返回下标的位置 查找不到返回-1
	// 6.1 strings.Index 从左往右出现的位置
	// str3 := "this中国你好goods你好golang"
	// str4 := "你好"
	// flag := strings.Index(str3, str4)
	// fmt.Println(flag) // 10 中文站3个字节

	// 6.2 strings.LastIndex 从右往左出现的位置
	str3 := "g中国你好thisgoods你好golang"
	str4 := "你好"
	flag := strings.LastIndex(str3, str4)
	fmt.Println(flag) // 22 中文站3个字节
}
