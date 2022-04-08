package main

import (
	"fmt"
)

// 在特殊情况下，需要在一个方法里获取不同的管道信息 可以使用select多路复用
// select 多路复用
func main() {
	// 定义 ch1管道并写入数据
	ch1 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch1 <- i
		fmt.Printf("【写入】数据ch1-%v成功\n", i)
	}

	// 定义 ch2管道并写入数据
	ch2 := make(chan string, 10)
	for i := 0; i < 10; i++ {
		ch2 <- fmt.Sprintf("数据_%d", i)
		fmt.Printf("【写入】数据ch2-%v成功\n", i)
	}

	// 使用select来获取chan的时候 不需要关闭close 管道
	// 并行读取ch1 ch2数据 ，与携程类似
	for {
		select {
		case v := <-ch1:
			fmt.Println("从ch1读取读取数据", v)
		case v := <-ch2:
			fmt.Println("从ch1读取读取数据", v)
		default:
			fmt.Println("读取完毕")
			return // 注意退出，否则是死循环
		}
	}

}
