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
	fmt.Printf("姓名%s 年龄%d 性别%s\n", p.Name, p.Age, p.Sex)
}

func (p Person) Test(a, b int, c string) {
	fmt.Printf("%d %d %s\n", a, b, c)
}

// 通过反射掉用方法
func main() {
	/*
		通过反射来进行方法掉用
		step1: 接口变量--> 对象反射对象:Value
		step2: 获取对应的方法对象：MethodByName()
		step3: 将方法对象进行掉用: Call()
	*/
	p1 := Person{"张三", 20, "男"}
	value := reflect.ValueOf(p1)
	fmt.Printf("Kind:%s ,type:%s\n", value.Kind(), value.Type()) // Kind:struct ,type:main.Person

	methodValue1 := value.MethodByName("PersionInfo")
	fmt.Printf("Kind:%s, type:%s\n", methodValue1.Kind(), methodValue1.Type()) // Kind:func, type:func()

	// 没有参数，进行掉用
	methodValue1.Call(nil) // 没有参数直接写nil
	args1 := make([]reflect.Value, 0)
	methodValue1.Call(args1) // 也可以声明一个空的切片也可以

	// 有参数进行掉用
	methodValue2 := value.MethodByName("Say")
	fmt.Printf("Kind:%s, type:%s\n", methodValue2.Kind(), methodValue2.Type()) // Kind:func, type:func(string)
	args2 := []reflect.Value{reflect.ValueOf("反射对象")}
	methodValue2.Call(args2) // hello 反射对象

	methodValue3 := value.MethodByName("Test")
	args3 := []reflect.Value{reflect.ValueOf(100), reflect.ValueOf(200), reflect.ValueOf("hello")}
	methodValue3.Call(args3) // 100 200 hello

}
