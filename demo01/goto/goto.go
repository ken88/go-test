package main

import "fmt"

func main() {

	fmt.Println("1")
	// 直接跳转到 gotoBreak 以后的代码执行 输出 1 4  ， 2 3不会执行
	goto gotoBreak

	fmt.Println("2")
	fmt.Println("3")

gotoBreak:
	fmt.Println(4)
}
