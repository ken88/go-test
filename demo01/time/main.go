package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取日期
	// getDate()

	// 获取当前时间戳
	// getTime()

	// 时间戳转日期
	// timeTurnDate()

	// 日期转时间戳
	// dateTurnTime()

	// 定时器
	// ticker()
	// timeObj := time.Now()

	// 程序执行时间
	executionTime()

}

// 程序执行时间
func executionTime() {

	startTime := time.Now().UnixNano() // 开始时间 纳秒
	/*执行的程序 */
	n := 5
	for {
		time.Sleep(time.Second)
		fmt.Println("aaa")
		n--
		if n == 0 {
			break
		}
	}
	endTime := time.Now().UnixNano() // 结束时间 纳秒

	seconds := float64((endTime - startTime) / 1e9) // 秒
	fmt.Printf("执行时间：%.2f秒\n", seconds)

	Milliseconds := float64((endTime - startTime) / 1e6) // 毫秒
	fmt.Printf("执行时间：%.2f毫秒\n", Milliseconds)

	nanoSeconds := float64(endTime - startTime) // 纳秒
	fmt.Printf("执行时间：%.2f纳秒\n", nanoSeconds)

}

// 毫秒定时器
func ticker() {

	// 1. 定时器
	// time.Now()
	// ticker := time.NewTicker(time.Second)
	// n := 5
	// for t := range ticker.C {
	// 	fmt.Println(t)
	// 	n--
	// 	if n == 0 {
	// 		ticker.Stop() // 停止定时器 销毁内存
	// 		break
	// 	}
	// }

	// 2. 休眠方法
	// fmt.Println("aaa")
	// time.Sleep(time.Second)
	// fmt.Println("bbb")
	// time.Sleep(time.Second)
	// fmt.Println("ccc")
	// time.Sleep(time.Second)
	// fmt.Println("ddd")

	// n := 5
	// for {
	// 	time.Sleep(time.Second)
	// 	fmt.Println("aaa")
	// 	n--
	// 	if n == 0 {
	// 		break
	// 	}
	// }
}

// 日期转时间戳
func dateTurnTime() {
	str := "2022-03-15 23:58:26" // 日期
	tmp := "2006-01-02 15:04:05" // 模版的格式
	timeObj, _ := time.ParseInLocation(tmp, str, time.Local)
	fmt.Println(timeObj)        // 2022-03-15 23:58:26 +0800 CST
	fmt.Println(timeObj.Unix()) // 1647359906
}

// 时间戳转日期
func timeTurnDate() {
	unixTime := 1647359906
	timeObj := time.Unix(int64(unixTime), 0) // 第一个参数是毫秒  ，第二个参数是纳秒
	str1 := timeObj.Format("2006-01-02 15:04:05")
	fmt.Println(str1)
}

// 获取当前时间戳
func getTime() {
	timeObj := time.Now()
	// 获取当前时间戳
	unixTime := timeObj.Unix() // 获取当前时间戳
	fmt.Println(unixTime)      // 1647359906
}

// 获取当前时间
func getDate() {
	timeObj := time.Now()
	fmt.Println(timeObj) // 2022-03-15 23:38:48.021147 +0800 CST m=+0.000096209

	// 1. 获取日期  年-月-日 时:分:秒
	year := timeObj.Year()     // 年
	month := timeObj.Month()   // 月
	day := timeObj.Day()       // 日
	hour := timeObj.Hour()     // 时
	minute := timeObj.Minute() // 分
	second := timeObj.Second() // 秒

	fmt.Printf("不够2位：%d-%d-%d %d:%d:%d \n", year, month, day, hour, minute, second) //2022-3-15 23:46:2
	// %02d 中的 2 表示宽度，如果整数不够2 列就补上0
	fmt.Printf("补位时间：%02d-%02d-%02d %02d:%02d:%02d \n", year, month, day, hour, minute, second) //2022-03-15 23:47:38

	// 2. 获取日期  年-月-日 时:分:秒
	// golang 里 2006=年  01=月 02=日 03=时（03表示12小时制 15 表示24小时制） 04=分 05=秒
	str := timeObj.Format("2006-01-02 03:04:05")
	fmt.Println("12小时", str) // 2022-03-16 12:03:17
	str1 := timeObj.Format("2006-01-02 15:04:05")
	fmt.Println("24小时", str1) // 2022-03-16 00:03:17
}
