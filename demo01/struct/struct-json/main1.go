package main

import (
	"encoding/json"
	"fmt"
)

// 学生结构体
type Student2 struct {
	Id     int
	Greder string
	Name   string
}

// 班级结构体
type Class struct {
	Title    string
	Students []Student2
}

func main() {
	var c = Class{
		Title:    "班级1",
		Students: make([]Student2, 0), // 给学生分配内存空间
	}
	// 给班级添加学生
	for i := 0; i < 10; i++ {
		var s = Student2{
			Id:     i,
			Greder: "男",
			Name:   fmt.Sprintf("str_%v", i),
		}
		c.Students = append(c.Students, s)
	}
	// fmt.Println(c)
	// 转json
	jsonByte, _ := json.Marshal(c)
	jsonStr := string(jsonByte)
	fmt.Println(jsonStr)
}
