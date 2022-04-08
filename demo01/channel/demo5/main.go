package main

import "fmt"

func main() {
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

	fmt.Println("主goroutine结束。。。")
}
