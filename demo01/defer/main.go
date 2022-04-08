package main

import "fmt"

// 下面这段代码输出的内容
func main() {
	defer_call()
	fmt.Println("结束")

	// fmt.Println(f1()) // 输出0
	// fmt.Println(f2()) // 输出100
	// fmt.Println(f3()) // 输出5
	// fmt.Println(f4()) // 输出0
	// fmt.Println(f5()) // 输出5
}

// 返回0 匿名返回值  返回值没有定义返回值的名字 就是匿名返回值
// 匿名返回值 return后在执行 defer
func f1() int {
	var a int
	defer func() {
		a++
	}()
	// 先return  在执行defer
	return a
}

// 返回100
// 命名返回值 ：
// 	1.返回值就是定义好返回值的本身 （定义返回值是a 返回的也是a） 执行完defer 在return
// 	2.返回值定义的不是定义好的返回值（定义返回值是a 返回值是其它变量） 先return 在defer
func f2() (a int) {
	// 先执行 defer 在return
	defer func() {
		a = 100
	}()
	// return a // 与直接返回return 是一样的
	return

}

// 返回5
// 返回值定义的不是定义好的返回值（定义返回值是y 实际返回值是其它变量） 先return 在defer
func f3() (y int) {
	x := 5
	defer func() {
		x = 100
	}()
	// 先return  在执行defer 返回值是x 不是y
	return x // 返回5
}

func f4() (x int) {
	// 先defer  在执行 return
	defer func(x int) {
		// 这里执行的是defer的x ，可以把x换成y不会影响外面的
		x++
		// fmt.Println("defer", x) // 1
	}(x)
	// fmt.Println("x=", x) // 0
	return x // 返回0
}

func f5() (x int) {
	defer func(x int) {
		// 这里执行的是defer的x 可以把x换成y 不会影响外面的
		x++
		fmt.Println("defer", x) // 1
	}(x)
	fmt.Println("x=", x) // 0
	return 5             // 返回5
}

// 执行顺序 后defer 先执行
func defer_call() {
	defer func() {
		fmt.Println("打印前")
	}()

	defer func() {
		fmt.Println("打印中")
	}()

	defer func() {
		// 不用recover() 接受异常，会报错
		// err := recover()
		// if err == nil {
		// 	fmt.Println(err)
		// }
		fmt.Println("打印后")
	}()

	panic("触发异常")
	// fmt.Println("我先执行") // 不会执行
	/*
		看下答案，输出：
		打印后
		打印中
		打印前
		panic: 触发异常

		参考解析：defer 的执行顺序是后进先出。当出现 panic 语句的时候，会先按照 defer 的后进先出的顺序执行，最后才会执行panic
	*/
}
