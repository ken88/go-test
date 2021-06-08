package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 随机数
func main() {
	// num1 := rand.Int()
	// fmt.Println(num1)
	// for i := 0; i < 10; i++ {
	// 	num2 := rand.Intn(90) // 0,100
	// 	fmt.Println(num2)
	// }
	fmt.Println(rand.Intn(10)) //无论打印错少次 都是一个数

	// 设置种子数，否则 每次rand.Intn(10) 都是一个值
	rand.Seed(time.Now().UnixNano())
	// for i := 0; i < 10; i++ {
	// 	num2 := rand.Intn(90) // 0,100
	// 	fmt.Println(num2)
	// }
	fmt.Println(rand.Intn(10)) // 设置了种子数，每次数字都不一样 ，随机输出小于10 的数
}
