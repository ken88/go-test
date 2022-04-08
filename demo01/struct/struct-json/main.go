package main

import (
	"encoding/json"
	"fmt"
)

// 首字母小写不能被json访问 首字母小写属于私有的
type Student struct {
	Id    int    `json:"id"`
	Green string `json:"green"`
	Name  string `json:"name"`
	Sno   string `json:"sno"`
}

type Student1 struct {
	Id    int
	Green string
	Name  string
	Sno   string
}

func main() {
	//  结构体转json 结构体没有定义标签 json的首字母都是大写的
	// structTojson()

	//  结构体转json 结构体定义了标签 json的首字母都是小写的
	structTojson1()

	// json转结构体
	// jsonToStruct()

}

// json转结构体 json首字母大写
func jsonToStruct() {
	var s1 Student
	//json转结构体
	strJson := `{"Id":12,"Green":"男","Name":"张三","Sno":"s001"}`

	err := json.Unmarshal([]byte(strJson), &s1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", s1) // main.Student{Id:12, Green:"男", Name:"张三", Sno:"s001"}
}

// json转结构体 json首字母小写
func structTojson1() {
	var s1 = Student{
		Id:    12,
		Green: "男",
		Name:  "张三",
		Sno:   "s001",
	}
	fmt.Printf("%#v \n", s1)        // main.Student{Id:12, Green:"男", Name:"张三", Sno:"s001"}
	jsonByte, _ := json.Marshal(s1) // 返回byte类型的切片
	jsonStr := string(jsonByte)     // 强制转换成string

	// 这里的json首字母都是大写的
	fmt.Println(jsonStr) // {"id":12,"green":"男","name":"张三","sno":"s001"}
}

// 结构体转json
func structTojson() {
	var s1 = Student1{
		Id:    12,
		Green: "男",
		Name:  "张三",
		Sno:   "s001",
	}
	fmt.Printf("%#v \n", s1)        // main.Student{Id:12, Green:"男", Name:"张三", Sno:"s001"}
	jsonByte, _ := json.Marshal(s1) // 返回byte类型的切片
	jsonStr := string(jsonByte)     // 强制转换成string

	// 这里的json首字母都是大写的
	fmt.Println(jsonStr) // {"Id":12,"Green":"男","Name":"张三","Sno":"s001"}
}
