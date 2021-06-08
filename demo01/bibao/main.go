package main

import "fmt"

func main() {
	// 1.调用increment()
	s1 := increment()
	fmt.Println(s1) // 0xe753a0  返回的increment（）方法里的 fun闭包地址

	// 2.调用increment（） 里返回的闭包方法
	res1 := s1()      // 调用increment（） 里返回的闭包方法
	fmt.Println(res1) // 输出 1

	// 3.再次调用increment（） 里返回的闭包方法
	res2 := s1()
	fmt.Println(res2) // 输出 2   每次执行都是执行 increment（）方法里的闭包

	// 3.重新调用increment（）
	s2 := increment()
	fmt.Println(s2) // 0xe753a0 返回的increment（）方法里的 fun闭包地址 与s1地址指向同一个地址
	res3 := s2()
	fmt.Println(res3) // 1  返回1  重新调用increment（） 执行返回的闭包方法重新开启另外一个地址

	res4 := s1()
	fmt.Println(res4) // 3 调用s1 执行increment()返回的方法，会继续按照s1里的匿名函数执行

	res5 := s2()
	fmt.Println(res5) //2  调用s2 执行increment()返回的方法，会继续按照s2里的匿名函数执行

}

func increment() func() int { // 外层函数

	// 1. 定义一个变量
	i := 0

	// 2. 定义一个匿名函数，给变量自增并返回
	fun := func() int {
		i++
		return i
	}

	// 3 返回该匿名函数
	return fun
}
