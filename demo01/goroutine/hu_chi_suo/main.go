package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup // 阻塞组

var mutex sync.Mutex // 声明互斥锁

// 全局变量表示票
var ticket = 10 // 10张票

// 临界资源测试
func main() {

	// 5个goroutine ,模拟5个售票口
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go saleTickets("售票窗口"+strconv.Itoa(i), i)
	}
	wg.Wait()
	fmt.Println("执行完毕")
}

func saleTickets(name string, i int) {

	rand.Seed(time.Now().UnixNano())
	for {
		mutex.Lock() // 加锁 一个携程执行完后 才可以执行下一个携程
		if ticket > 0 {
			// 进入到睡眠
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)

			fmt.Println("携程", i, name, "售出", ticket)
			ticket--
		} else {
			mutex.Unlock() // 最后一张票售出，解锁当前goroutine
			fmt.Println("携程", i, name, "票售馨,没有票了")
			break
		}
		mutex.Unlock() // 解锁
	}

	wg.Done()
}
