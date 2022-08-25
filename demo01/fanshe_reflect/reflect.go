package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 反射操作：通过反射，获取一个借口类型的变量的类型和数值
	var x float64 = 3.14
	fmt.Println("type", reflect.TypeOf(x))   // type float64
	fmt.Println("value", reflect.ValueOf(x)) // value 3.14

	fmt.Println("------------------------------")
	v := reflect.ValueOf(x)
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64) // kind is float64: true
	fmt.Println("type:", v.Type())                               // type: float64
	fmt.Println("value:", v.Float())                             // value: 3.14
}
