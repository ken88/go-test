package main

import "fmt"

func main() {
	//1. a的变量指向一个内存地址
	// a := 10
	// b := &a
	// // 所有的变量都有自己的内存地址
	// fmt.Printf("a地址：%v a的类型：%T a的地址：%p \n", a, a, &a) // a地址：10 a的类型：int a的地址：0xc0000160b8
	// fmt.Printf("b地址：%v b的类型：%T b的地址：%p \n", b, b, &b) // b地址：0xc0000160b8 b的类型：*int b的地址：0xc00000e028
	// fmt.Println("通过b打印a的值", *b)                       // 通过b打印a的值 10

	//2. new方法给指针分配内存空间  返回的是指针类型 make 分配 slice map channel  实际项目中用不到
	// var a *int
	// a = new(int)
	// fmt.Printf("值%v 类型%T 地址%p\n", a, a, &a) // 值0xc0000160b8 类型*int 地址0xc00000e028
	// *a = 100
	// fmt.Printf("值%v 类型%T 地址%p\n", a, a, &a)  // 值0xc0000160b8 类型*int 地址0xc00000e028
	// fmt.Printf("值%v 类型%T 地址%p\n", *a, a, &a) // 值100 类型*int 地址0xc00000e028

	// 3. 测试改变内存地址数据
	// f1()

}

func f1() {
	a := 5
	f2(a)
	fmt.Println("a=", a) // 5

	f3(&a)
	fmt.Println("a=", a) // 40

}

func f2(x int) {
	x = 10
}

func f3(x *int) {
	*x = 40
}
