package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	f1 := fun1
	value := reflect.ValueOf(f1)
	fmt.Printf("Kind:%s type:%v\n", value.Kind(), value.Type()) // Kind:func type:func()

	value2 := reflect.ValueOf(fun2)
	fmt.Printf("Kind:%s type:%v\n", value2.Kind(), value2.Type()) // Kind:func type:func(int, string)

	value3 := reflect.ValueOf(fun3)
	fmt.Printf("Kind:%s type:%v\n", value3.Kind(), value3.Type()) // Kind:func type:func(int, string) string

	// 通过反射进行函数调用
	value.Call(nil)

	// 有参数的
	value2.Call([]reflect.Value{reflect.ValueOf(100), reflect.ValueOf("hello")})

	// 有参数并且有返回值的
	resuleValue := value3.Call([]reflect.Value{reflect.ValueOf(200), reflect.ValueOf("张三")})
	fmt.Println(resuleValue)                                                      // [张三200]
	fmt.Println(len(resuleValue))                                                 // 1
	fmt.Printf("Kind:%s type:%v\n", resuleValue[0].Kind(), resuleValue[0].Type()) // Kind:string type:string
	s := resuleValue[0].Interface().(string)
	fmt.Println("返回数据:", s)    // 返回数据: 张三200
	fmt.Printf("返回类型：%T\n", s) // 返回类型：string

}

func fun1() {
	fmt.Println("我是fun1(),无参的...")
}

func fun2(i int, s string) {
	fmt.Println("我是fun2(),有参的...", i, s)
}

func fun3(i int, s string) string {
	fmt.Println("fun3(),有参的,也有返回值", i, s)
	return s + strconv.Itoa(i)
}
