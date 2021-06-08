package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // 协程

func main() {
	var ch = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	for i := 1; i <= 10; i++ {
		fmt.Println(<-ch)
	}
}

// 通过for 打印管道信息
func for1() {
	var ch = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	// 用for打印管道 不需要关闭管道
	for i := 1; i <= 10; i++ {
		fmt.Println(<-ch)
	}
}

// 用for range 打印管道
func forRange() {
	var ch = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
	// for range 循环遍历管道的值，注意：管道没有key
	//用for range 循环 要关闭管道 否则会出现管道阻塞 fatal error: all goroutines are asleep - deadlock!
	for v := range ch {
		fmt.Println(v)
	}
}

func test1() {
	// 1. 创建channel
	ch := make(chan int, 3)

	// 2. 给管道里面存储数据
	ch <- 10
	ch <- 20
	ch <- 30

	// 3. 获取管道里面的内容
	a := <-ch
	fmt.Println(a) // 10
	<-ch           // 从管道取出20  // 20

	b := <-ch
	fmt.Println(b) //30

	// 管道阻塞
	ch1 := make(chan int, 1)
	ch1 <- 11
	ch1 <- 12 // all goroutines are asleep - deadlock!  管道阻塞
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
