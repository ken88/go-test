package main

import "fmt"

// 自定义类型
type myInt int

// 定义类型别名
type myInt1 = int

func main() {
	var a myInt = 10
	fmt.Printf("a的值%v 类型%T\n", a, a) // a的值10 类型main.myInt

	// 类型定义别名 类型还是原始类型
	var b myInt1 = 20
	fmt.Printf("b的值%v 类型%T\n", b, b) // b的值20 类型int
}
