package main

import (
	"fmt"
)

func main() {

	ch := make(chan int) //这里就是创建了一个channel，这是无缓冲管道注意,容量为0的通道
	go func() {
		fmt.Println("子goroutine开始执行")
		// time.Sleep(3 * time.Second)
		data := <-ch
		fmt.Println("data", data)
	}()
	fmt.Println("主goroutine开始执行")
	ch <- 10

	// test1()
	fmt.Println("主goroutine结束。。。")
}

func test1() {
	ch := make(chan int) //这里就是创建了一个channel，这是无缓冲管道注意,容量为0的通道
	go func() {          //创建子go程
		for i := 0; i < 4; i++ {
			ch <- i //循环写入管道
			fmt.Println("写入", i)
		}
	}()

	for i := 0; i < 4; i++ { //主go程
		num := <-ch //循环读出管道
		fmt.Println("读出", num)
	}

}
