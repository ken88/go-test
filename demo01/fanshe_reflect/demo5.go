package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name   string
	Age    int
	School string
}

// 结构体通过反射修改值
func main() {
	s1 := Student{"张三", 12, "xx教育"}

	// 通过反射，更改对象的数值：前提也是数据可以修改
	fmt.Printf("%T\n", s1) // main.Student
	fmt.Println(s1)        // {张三 12 xx教育}

	p1 := &s1
	fmt.Printf("%T\n", p1) // *main.Student
	fmt.Println(s1.Name)
	fmt.Println((*p1).Name, p1.Name)

	// 改变数值
	value := reflect.ValueOf(&s1)
	if value.Kind() == reflect.Ptr {
		newValue := value.Elem()
		fmt.Println("是否可以修改值", newValue.CanSet()) // 是否可以修改值 true

		f1 := newValue.FieldByName("Name")
		f1.SetString("李四")
		fmt.Println(f1) // 李四

		f3 := newValue.FieldByName("School")
		f3.SetString("幼儿园")
		fmt.Println(f3) // 幼儿园

		fmt.Println(s1) // 再次打印s1 变成了 {李四 12 幼儿园}
	}
}
