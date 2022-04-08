package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 边写边读 goroutine-channel
func main() {
	var ch = make(chan int, 10)
	// 1. 开启写管道携程
	wg.Add(1)
	go f1(ch)

	// 2. 开启读管道携程
	wg.Add(1)
	go f2(ch)

	// 等待携程
	wg.Wait()

	fmt.Println("执行完成1")
}

// 写管道
func f1(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("【写入】数据%v成功\n", i)
		time.Sleep(time.Microsecond * 10)
	}
	close(ch)
	wg.Done()
}

// 读管道
func f2(ch chan int) {
	for v := range ch {
		fmt.Printf("【读取】数据%v成功\n", v)
	}
	wg.Done()
}
