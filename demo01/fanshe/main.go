package main

import (
	"fmt"
	"reflect"
)

type myInt int
type Persion struct {
	Name string
	Age  int
}

func main() {
	a := 10
	b := 12.2
	c := true
	d := "hello golang"
	var e myInt
	var f = Persion{
		Name: "张三",
		Age:  20,
	}
	e = 20
	reflectFn(a) // 类型：int 类型名称：int 类型种类：int
	reflectFn(b) // float64 类型名称：float64 类型种类：float64
	reflectFn(c) // bool 类型名称：bool 类型种类：bool
	reflectFn(d) // string 类型名称：string 类型种类：string
	reflectFn(e) // main.myInt 类型名称：myInt 类型种类：int
	reflectFn(f) // main.Persion 类型名称：Persion 类型种类：struct
	var g = 29
	reflectFn(&g) // *int 类型名称： 类型种类：ptr
	h := [3]int{1, 2, 3}
	reflectFn(h) // 类型：[3]int 类型名称： 类型种类：array
	i := []int{4, 5, 6}
	reflectFn(i) // []int 类型名称： 类型种类：slice

}

// 查看类型
func reflectFn(x interface{}) {
	v := reflect.TypeOf(x)
	// v.Name() // 类型名称
	// v.Kind() // 类型种类
	fmt.Printf("类型：%v 类型名称：%v 类型种类：%v\n", v, v.Name(), v.Kind())
}
