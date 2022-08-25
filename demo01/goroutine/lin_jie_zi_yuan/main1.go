package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// 全局变量表示票
var ticket = 10 // 10张票

// 临界资源测试
func main() {

	// 5个goroutine ,模拟5个售票口
	for i := 1; i <= 10; i++ {
		go saleTickets("售票窗口" + strconv.Itoa(i))
	}
	time.Sleep(time.Second * 5)
}

func saleTickets(name string) {
	rand.Seed(time.Now().UnixNano())
	for {
		if ticket > 0 {
			// 进入到睡眠
			/*
				逻辑分析：
				5个窗口 ，都进入到这里，开始睡秒 g1 g2 g3 g4 g5
				例如：g3 睡眠时间到 可能是最后一张票了，票数-1  当前票数0 退出循环 g3 goroutine销毁
				     这个时候，g1 携程醒了 也减掉了一张票 这个时候 票数就成负数了  g1 goroutine销毁
					 这个时候，g3 携程醒了 也减掉了一张票 这个时候 票数一值是 （-N+1）  g3 goroutine销毁
					 直到所有的携程都执行完
			*/
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)

			fmt.Println(name, "售出", ticket)
			ticket--
		} else {
			fmt.Println(name, "票售馨,没有票了")
			break
		}
	}
}
