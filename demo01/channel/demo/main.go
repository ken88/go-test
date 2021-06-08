package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	start := time.Now().Unix()

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 50000)
	exitChan := make(chan bool, 16) //标识primeChan close

	//存放数字的协程
	wg.Add(1)
	go putNum(intChan)

	//统计素数的协程
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go primeNum(intChan, primeChan, exitChan)
	}

	//打印素数的协程
	wg.Add(1)
	go printPrime(primeChan)

	//判断exitChan是否存满值
	wg.Add(1)
	go func() {
		for i := 0; i < 16; i++ {
			<-exitChan
		}
		//关闭primeChan
		close(primeChan)
		wg.Done()
	}()

	wg.Wait()

	end := time.Now().Unix()
	fmt.Println("执行完毕....", end-start, "毫秒")

}

func putNum(intChan chan int) {
	for i := 2; i < 120000; i++ {
		intChan <- i
	}
	close(intChan)
	wg.Done()
}

// 从intChan取出数据，并判断是否为素数，是素数放到就把intChan放到primeNum里
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for num := range intChan {
		var flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num // num是素数
		}
	}
	// close(primeChan) // 如果一个chanel关闭了，就没发给这个channel发送数据了

	exitChan <- true
	wg.Done()
}

// 打印素数
func printPrime(primeChan chan int) {
	// for v := range primeChan {
	// 	fmt.Println(v)
	// }
	wg.Done()
}
