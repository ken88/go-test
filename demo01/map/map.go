package main

import "fmt"

func main() {
	//第一种创建
	var userinfo = make(map[string]string)
	userinfo["username"] = "张三"
	userinfo["age"] = "20"
	userinfo["sex"] = "男"
	fmt.Println(userinfo)             //map[age:20 sex:男 username:张三]
	fmt.Println(userinfo["username"]) //张三

	//第二种创建
	// var userinfo = map[string]string{
	// 	"username": "张三",
	// 	"age":      "20",
	// 	"sex":      "男",
	// }
	// fmt.Println(userinfo)             //map[age:20 sex:男 username:张三]
	// fmt.Println(userinfo["username"]) //张三

	//第三种创建
	// userinfo := map[string]string{
	// 	"username": "张三",
	// 	"age":      "20",
	// 	"sex":      "男",
	// }
	// userinfo["username"] = "李四"
	// userinfo["age"] = "40"
	// // v, ok := userinfo["xxx"] //获取map里是否包含的key ok有返回true v返回当前值 没有返回false
	// // fmt.Println(v, ok)
	// delete(userinfo, "username") //删除
	// fmt.Println(userinfo)
	// for key, val := range userinfo {
	// 	fmt.Printf("key:%v val:%v\n", key, val)
	// }

	//定义元素类型为map类型的切片
	// var userinfo = make(map[string][]string)
	// userinfo["hobby"] = []string{
	// 	"吃饭",
	// 	"睡觉",
	// }
	// userinfo["work"] = []string{
	// 	"php",
	// 	"golang",
	// }
	// fmt.Println(userinfo)
	// for _, val := range userinfo {
	// 	for _, value := range val {
	// 		fmt.Println(value)
	// 	}

	// }
}
