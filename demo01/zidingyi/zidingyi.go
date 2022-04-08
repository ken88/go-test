package zidingyi

import "time"

// 返回当前时间
func GetDate() string {

	timeObj := time.Now()

	// golang 里 2006=年  01=月 02=日 03=时（03表示12小时制 15 表示24小时制） 04=分 05=秒
	return timeObj.Format("2006-01-02 03:04:05")
}
