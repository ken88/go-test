package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
	Sex  string
}

func (p Person) Say(msg string) {
	fmt.Println("hello", msg)
}

func (p Person) PersionInfo() {
	fmt.Printf("姓名%s 年龄%d 性别%s", p.Name, p.Age, p.Sex)
}

func main() {
	// 未知原有类型
	p1 := Person{"张三", 30, "男"}
	GetMessage(p1)
}

// 获取input的信息
func GetMessage(input interface{}) {
	getType := reflect.TypeOf(input)           // 先获取input的类型
	fmt.Println("get type is", getType.Name()) // get type is Person
	fmt.Println("get Kind is", getType.Kind()) // get Kind is struct

	// 获取数值
	getValue := reflect.ValueOf(input)
	fmt.Println(getValue) // {张三 30 男}

	// 获取字段
	/*
		step1:先获取type对象:reflect.type
			Numfield()
			Field(index)
		step2:通过Field获取每一个Field字段
		step3:通过Interface(),得到对应的value
	*/
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface() // 获取第一个数值
		fmt.Printf("字段的名称%v 字段的类型%v 字段的值%v\n", field.Name, field.Type, value)
	}

	// 获取方法
	for i := 0; i < getType.NumMethod(); i++ {
		method := getType.Method(i)
		fmt.Printf("方法名称%v,方法类型%v\n", method.Name, method.Type)
	}

}
