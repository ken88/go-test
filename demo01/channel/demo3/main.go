package main

import "fmt"

func main() {

	// 可读可写管道（双向管道）
	// writeAndRead()

	// 指写管道 不能读（单项管道）
	// writeChan()

	//  指读管道 不能写（单项管道）
	// readChan()
}

// 可读可写管道（双向管道）
func writeAndRead() {
	ch := make(chan int, 2)
	ch <- 10
	ch <- 20
	fmt.Println(<-ch)
}

// 只写管道 chan<-  不能读（单项管道）
func writeChan() {
	ch := make(chan<- int, 2)
	ch <- 10
	ch <- 20
	// fmt.Println(<-ch) // invalid operation: <-ch (receive from send-only type chan<- int)
}

// 指读管道  <-chan 不能写（单项管道）
func readChan() {
	// ch := make(<-chan int, 2)
	// ch <- 10 // 不能写入 invalid operation: cannot send to receive-only
}
