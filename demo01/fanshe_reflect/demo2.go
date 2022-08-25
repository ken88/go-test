package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 已知原有类型
	var num float64 = 3.14

	// 接口类型变量-> 反射类型变量
	value := reflect.ValueOf(num)
	fmt.Println(value) // 3.14

	// 反射类型对象-> 接口类型变量
	converValue := value.Interface().(float64)
	fmt.Println(converValue) //  3.14

	//  反射类型对象-> 接口类型变量 理解为强制转换
	// Golang对类型要求非常严格，类型一定要完全符合
	// 一个是*float64 ，一个是float64 ，如果弄混，直接panic
	pointer := reflect.ValueOf(&num)
	fmt.Println(pointer) // 0x140000160c8
	// converPointer := pointer.Interface().(float64) // panic: interface conversion: interface {} is *float64, not float64
	converPointer := pointer.Interface().(*float64)

	fmt.Println(converPointer) // 0x140000160c8
}
