package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var ch = make(chan int, 10)

func main() {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(<-ch)
	// }
	for v := range ch {
		fmt.Println(v)
	}

}

func test() {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	for i := 1; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

func test2() {
	for i := 11; i <= 20; i++ {
		ch <- i
	}
	fmt.Println("test2..................")
	close(ch)
	wg.Done()
}

func printChannel() {
	for {
		fmt.Println("管道输出等待中！")
		i, isClose := <-ch
		if !isClose {
			fmt.Println("管道已经关闭！")
			// ch <- 123
			break
		} else {
			fmt.Println("管道是否开启：", isClose, "----值：", i)
		}

	}
	wg.Done()
}
