package main

import (
	"fmt"

	"io/ioutil"
	"net/http"
)

func main() {

	for i := 0; i < 30; i++ {
		url := "http://101.132.109.207:8000/logistics/add-trackingmore"
		req, _:= http.NewRequest("GET", url, nil)
		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
		fmt.Printf("第%d次请求结束\n", i)
	}
}

// 结构体转换json 与json转换对象
// func structjson() {
// 	// 1. 对象转json
// 	var p1 = Person1{
// 		Name: "张三",
// 		Age:  20,
// 		Sex:  "男",
// 	}

// 	// fmt.Printf("%#v %T\n", p1, p1) // main.Person{name:"张三", age:20, sex:"男"} main.Person

// 	// 将p1对象转换为Byte类型的切片
// 	jsonByte, _ := json.Marshal(p1)
// 	// fmt.Printf("%#v %T\n", jsonByte, jsonByte) //[]byte{0x7b, 0x7d} []uint8
// 	// 将切片转换json
// 	jsonStr := string(jsonByte)
// 	fmt.Printf("%v\n", jsonStr) // {"Name":"张三","Age":20,"Sex":"男"}

// 	// 2. json 转对象
// 	var str1 = `{"Name":"张三","Age":20,"Sex":"男"}`
// 	var s2 Person1
// 	err := json.Unmarshal([]byte(str1), &s2)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Printf("%#v\n", s2) //main.Person1{Name:"张三", Age:20, Sex:"男"}
// }

//结构体
// func structs() {

// 	// //Person P大写结构体公有的 小写私有的
// 	type Person struct {
// 		name string // 小写为私有的
// 		age  int
// 		sex  string
// 	}

// 	// //大写为公共 才可以转换json  私有属性不能被json包访问
// 	// type Person1 struct {
// 	// 	Name string
// 	// 	Age  int
// 	// 	Sex  string `json:"xxx"` // 增加tag标签 转换后{"Name":"张三","Age":20,"xxx":"男"} 把sex转换xxx 名字自定义
// 	// }

// 	//结构体初始化

// 	//第1种初始化方式
// 	var p1 Person
// 	p1.name = "张三"
// 	p1.age = 20
// 	p1.sex = "男"
// 	fmt.Printf("%#v %T\n", p1, p1)      // main.Person{name:"张三", age:20, sex:"男"} main.Person
// 	fmt.Printf("%#v %T\n", p1.name, p1) // "张三" main.Person

// 	//第2种初始化方式
// 	var p2 = new(Person)
// 	p2.name = "李四"
// 	p2.age = 30
// 	p2.sex = "女"
// 	(*p2).sex = "女"
// 	fmt.Printf("%#v %T\n", p2, p2)      // &main.Person{name:"李四", age:30, sex:"女"} *main.Person
// 	fmt.Printf("%#v %T\n", p2.name, p2) // "李四" *main.Person

// 	//第3种初始化方式
// 	var p3 = &Person{}
// 	p3.name = "赵四"
// 	p3.age = 22
// 	p3.sex = "男"
// 	fmt.Printf("%#v %T\n", p3, p3) // &main.Person{name:"赵四", age:22, sex:"男"} *main.Person

// 	//第4种初始化方式
// 	var p4 = Person{
// 		name: "哈哈",
// 		age:  25,
// 		sex:  "女",
// 	}
// 	fmt.Printf("%#v %T\n", p4, p4) // main.Person{name:"哈哈", age:25, sex:"女"} main.Person

// 	//第5种初始化方式
// 	var p5 = &Person{
// 		name: "王麻子",
// 		age:  30,
// 		sex:  "男",
// 	}
// 	fmt.Printf("%#v %T\n", p5, p5) // &main.Person{name:"王麻子", age:30, sex:"男"} *main.Person

// 	//第6种初始化方式
// 	var p6 = &Person{
// 		name: "旺旺",
// 	}
// 	fmt.Printf("%#v %T\n", p6, p6) //&main.Person{name:"旺旺", age:0, sex:""} *main.Person

// 	//第7种初始化方式
// 	var p7 = &Person{
// 		"张三",
// 		30,
// 		"男",
// 	}
// 	fmt.Printf("%#v %T\n", p7, p7) // &main.Person{name:"张三", age:30, sex:"男"} *main.Person
// }

//定时器
// func tickers() {
// 	//定时器 ticker
// 	ticker := time.NewTicker(time.Second) //time.Second 一秒钟执行

// 	n := 5
// 	//ticker.C
// 	for t := range ticker.C {
// 		n--
// 		if n == 0 {
// 			ticker.Stop() //中止定时器
// 			break
// 		}
// 		fmt.Println(t)
// 	}

// 	//time.Sleep("aa") //休眠方法 1秒钟
// }

// //日期
// func times() {
// 	// 获取日期
// 	timeObj := time.Now() // 获取时间  2020-09-14 15:38:02.3043408 +0800 CST m=+0.010970101
// 	fmt.Println(timeObj)
// 	year := timeObj.Year()
// 	month := timeObj.Month()
// 	day := timeObj.Day()
// 	hour := timeObj.Hour()
// 	minute := timeObj.Minute()
// 	second := timeObj.Second()
// 	fmt.Printf("%d-%d-%d %d:%d:%d\n", year, month, day, hour, minute, second)           //2020-9-14 15:41:32
// 	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second) //输出值有一位的话，补0 2020-09-14 15:42:34
// 	fmt.Println(timeObj)
// 	//03代表12小时制  15代表 24小时制
// 	str := timeObj.Format("2006-01-02 03:04:05")
// 	fmt.Println("12小时制：", str)
// 	str1 := timeObj.Format("2006-01-02 15:04:05")
// 	fmt.Println("24小时制：", str1)

// 	// 获取时间戳
// 	unixtime := timeObj.Unix() //获取当前的时间戳 (毫秒)
// 	fmt.Println("当前时间戳：", unixtime)

// 	// 时间戳转换日期
// 	unixtime1 := 1600069874 //2020-09-14 03:51:14
// 	dates := time.Unix(int64(unixtime1), 0)
// 	str2 := dates.Format("2006-01-02 15:04:05")
// 	fmt.Println("时间戳转换时间：", str2) // 2020-09-14 15:51:14

// 	// 日期字符串 转换时间戳
// 	str3 := "2020-09-14 15:51:14"
// 	var tmp = "2006-01-02 15:04:05"
// 	res, _ := time.ParseInLocation(tmp, str3, time.Local)
// 	fmt.Println(res)        //2020-09-14 15:51:14 +0800 CST
// 	fmt.Println(res.Unix()) //1600069874
// }

// // map 排序
// func sortMapAsc(sorts map[string]string) string {
// 	//main
// 	// m1 := map[string]string{
// 	// 	"username": "张三",
// 	// 	"age":      "20",
// 	// 	"sex":      "男",
// 	// 	"height":   "180",
// 	// }

// 	// res := sortMapAsc(m1)
// 	// fmt.Println(res)

// 	var sliceKey []string
// 	for k, _ := range sorts {
// 		sliceKey = append(sliceKey, k)
// 	}
// 	sort.Strings(sliceKey)
// 	var res string
// 	for _, val := range sliceKey {
// 		res += fmt.Sprintf("%v=>%v", val, sorts[val])
// 	}
// 	return res
// }

// // 多参数 （...代表可变参数）
// func sums(x string, y ...int) string {

// 	//放到主函数里执行
// 	// sums := sums("返回值", 1, 2, 3, 4, 5)
// 	// fmt.Println(sums)

// 	sum := 0
// 	for _, val := range y {
// 		sum += val
// 	}
// 	res := x + strconv.Itoa(sum) //讲返回值转换为string类型

// 	return res
// }
