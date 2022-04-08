package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // 协程
var count = 10

var mutex sync.Mutex // 声明互斥锁

// 通过 go build --race main.go 编译后查看
// 互斥锁
func main() {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go test()
	}
	wg.Wait()
}

func test() {
	mutex.Lock() // 加锁
	if count > 0 {
		time.Sleep(2)
		count--
		fmt.Println("this count is", count)

	}
	mutex.Unlock() //解锁
	wg.Done()
}
