package main

import (
	"fmt"
	"reflect"
)

// 通过反射修改值
func main() {
	num := 3.14
	fmt.Println("num=", num)

	// 需要操作指针
	pointer := reflect.ValueOf(&num) // 注意，参数必须是指针，才可以修改值，如果不传指针，会出现恐慌
	newValue := pointer.Elem()

	fmt.Println("类型:", newValue.Type())       // 类型: float64
	fmt.Println("是否可以修改值", newValue.CanSet()) // 是否可以修改值 true
	newValue.SetFloat(3.111)                  // 修改新值
	fmt.Println("num=", num)                  // num= 3.111

	// 如果reflect.ValueOf的参数不是指针?
	// value := reflect.ValueOf(num)
	// value.SetFloat(1.12) // panic: reflect: reflect.Value.SetFloat using unaddressable value
	// fmt.Println(value.CanSet()) // false
	// value.Elem() // 如果非指针，会直接panic panic: reflect: call of reflect.Value.Elem on float64 Value
}
