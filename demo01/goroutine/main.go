package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // 协程

func test1() {
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println("test1 输出值", i)
	}
	wg.Done() // 协程计数器-1
}
func test2() {
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println("test2 输出值", i)
	}
	wg.Done() // 协程计数器-1
}

func test3(num int) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Printf("协程（%v）打印的第%v条数据\n", num, i)
	}
}
func main() {

	/*
		go test1() //开启协成
		go test2() //开启协成
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Millisecond * 20)
			fmt.Println("main 输出值", i)
		}
		fmt.Println("主进程退出")

			开启协成 主进程只要执行完毕，不管协成是否执行完成，结束当前进程
			打印结果
			main 输出值 1
			main 输出值 2
			main 输出值 3
			main 输出值 4
			test2 输出值 1
			test1 输出值 1
			main 输出值 5
			main 输出值 6
			main 输出值 7
			main 输出值 8
			test2 输出值 2
			test1 输出值 2
			main 输出值 9
			main 输出值 10
			主进程退出

	*/
	// cmpNum := runtime.NumCPU() // 获取当前计算器上面的cpu个数
	// fmt.Println("cupNum", cmpNum)

	// //可以自己设置使用多个cpu
	// runtime.GOMAXPROCS(cmpNum - 1)

	/*
		// 正常打印
			wg.Add(1)  // 协程计数器加1
			go test1() // 表示开启协程
			wg.Add(1)  // 协程计数器加1
			go test2() // 表示开启协程

			for i := 1; i <= 10; i++ {
				time.Sleep(time.Millisecond * 20)
				fmt.Println("main 输出值", i)
			}
			wg.Wait() //如果有协程没有执行完等待， 所有协程全部执行完 继续往下走
			fmt.Println("主进程退出")
	*/
	// for i := 1; i <= 6; i++ {
	// 	wg.Add(1)
	// 	go test3(i)
	// }
	// wg.Wait()
	// fmt.Println("主线程打印完毕")

	// 需求：统计1-12000的数字中，哪些是素数
	goroutines()
}
func test(n int) {
	fmt.Println("协程", n)
	for num := (n-1)*30000 + 1; num < n*30000; num++ {
		if num > 1 {
			var flag = true
			for i := 2; i < num; i++ {
				if num%i == 0 {
					flag = false
					break
				}
			}
			if flag {
				// fmt.Println("协程", n, "执行：", num, "数素数")
			}
		}
	}
	wg.Done()
}

// goroutine 找出1-120000的素数
func goroutines() {
	start := time.Now().Unix()
	fmt.Println("开启协程准备...")
	for i := 1; i <= 4; i++ {
		fmt.Println("开启协程", i)
		wg.Add(1)
		go test(i)
	}
	wg.Wait()
	end := time.Now().Unix()
	fmt.Println("执行时间：", end-start) // 5毫秒
}

func foreach() {
	// 通过for 找出1-120000的素数
	start := time.Now().Unix()
	for num := 2; num < 120000; num++ {
		var flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			// fmt.Println(num, "数素数")
		}
	}

	end := time.Now().Unix()
	fmt.Println("执行时间：", end-start) // 9毫秒
}
