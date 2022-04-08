package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // 协程

func main() {
	// 单条录入管道
	// test1()

	// 通过for打印chan值
	// printChan()

	//通过forRange打印chan值
	forRange()

	// 通过携程打印素数
	// goroutines()
}

// 通过for打印chan值
func printChan() {
	var ch = make(chan int, 10)
	// 增加数据到chan
	for i := 1; i <= 10; i++ {
		ch <- i
	}

	// 取出chan的信息
	for i := 1; i <= 10; i++ {
		fmt.Println(<-ch)
	}
}

// 用for range 打印管道
func forRange() {
	// 信息录入到chan
	var ch = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch) // 关闭管道

	// 打印管道的值
	// for range 循环遍历管道的值，注意：管道没有key
	// 用for range 循环 要关闭管道 否则会出现管道阻塞 fatal error: all goroutines are asleep - deadlock!
	for v := range ch {
		fmt.Println(v)
	}
}

// 单条录入管道
func test1() {
	// 1. 创建channel
	ch := make(chan int, 3)

	// 2. 给管道里面存储数据
	ch <- 10
	ch <- 20
	ch <- 30
	// 查看管道的容量
	fmt.Printf("值：%v 长度：%v 容量：%v\n", ch, len(ch), cap(ch)) // 值：0xc000020180 长度：3 容量：3

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
		// fmt.Println(num)
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

// 通过携程打印素数
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
