package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // 协程

var mutex sync.RWMutex // 声明读写互斥锁

// 通过 go build --race main.go 编译后查看
// 读写互斥锁
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
		wg.Add(1)
		go read()
	}
	wg.Wait()
}

// 写互斥锁，当一goroutine进行写的时候，其它goroutine 不能读 也不能写
func write() {
	mutex.Lock()
	fmt.Println("write执行写操作！")
	time.Sleep(time.Second)
	mutex.Unlock()
	wg.Done()
}

// 读不互斥
func read() {
	mutex.RLock()
	fmt.Println("read执行读操作！")
	mutex.RUnlock()
	wg.Done()
}
