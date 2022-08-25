package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // 协程

var rwMutex sync.RWMutex // 声明读写互斥锁

// 通过 go build --race main.go 编译后查看
// 读写互斥锁
func main() {
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go write(i)
		wg.Add(1)
		go read(i)
	}
	wg.Wait()
	fmt.Println("main 结束...")
}

// 写互斥锁，当一goroutine进行写的时候，其它goroutine 不能读 也不能写
func write(i int) {

	fmt.Println(i, "开始写 write start...")

	rwMutex.Lock() // 写操作上锁
	fmt.Println(i, "write正在写数据: reading...")
	time.Sleep(3 * time.Second)
	rwMutex.Unlock() // 写操作解锁
	fmt.Println(i, "写操作结束")

	defer wg.Done()
}

// 读不互斥  可以随便读，多个goroutine同时读
func read(i int) {

	fmt.Println(i, "开始读 read start...")

	rwMutex.RLock() // 读操作上锁
	fmt.Println(i, "read正在读数据: reading...")
	// time.Sleep(3 * time.Second)
	rwMutex.RUnlock() // 读操作解锁

	fmt.Println(i, "读操作结束")
	defer wg.Done()
}
