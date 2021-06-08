package main

import "fmt"

// 闭包
func fibonacci() func() int {
	a, b := 0, 1
	fmt.Printf("外层a的值%d\n", a)
	return func() int {
		a, b = b, a+b
		fmt.Printf("a的值%d\n", a)
		return a
	}
}
func main() {

	f := fibonacci()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}
