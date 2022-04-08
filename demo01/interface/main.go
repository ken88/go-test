package main

import (
	"fmt"
)

type Animaler interface {
	SetName(string)
	GetName() string
}
type Dog struct {
	Name string
}

func (d *Dog) SetName(name string) {
	d.Name = name
}
func (d Dog) GetName() string {
	return d.Name
}

type Address struct {
	Name  string
	phone string
}

func main() {

	// dog继承接口 实现方法
	// f1()

	// 接口类型切片
	// f2()

	// map类型接口
	f3()

	// 任意类型传递
	// f4()

	//类型断言
	// justifyType(123)

}

// 类型断言 x.(type) 只能用在switch里面
func justifyType(x interface{}) {

	switch v := x.(type) {
	case string:
		fmt.Println("这是string类型", v)

	case int:
		fmt.Println("这是int类型", v)

	case bool:
		fmt.Println("这是bool类型", v)

	default:
		fmt.Println("无结果", v)

	}
}

func f4() {
	// show(123) //123 int
	show("abc") // abc string
}

// 空接口可以是任意类型
func show(a interface{}) {
	fmt.Printf("%v %T\n", a, a)
}

// dog继承接口 实现方法
func f1() {

	// Dog实现Animaler接口
	d := &Dog{
		Name: "小黑",
	}
	var d1 Animaler = d
	fmt.Println(d1.GetName())
	d1.SetName("阿奇")
	fmt.Println(d1.GetName())
}

// 切片接口类型
func f2() {
	type A interface{}
	var a A
	var str = "你好golang"
	a = str                     //让字符串实现A这个接口
	fmt.Printf("%v %T\n", a, a) // 你好golang string

	a = 123
	fmt.Printf("%v %T\n", a, a) // 123 int

	a = true
	fmt.Printf("%v %T\n", a, a) // true bool

	var s1 = []interface{}{1, 2, '3', true}
	fmt.Println(s1)
}

// map类型接口
func f3() {
	userinfo := make(map[string]interface{})
	userinfo["username"] = "张三"
	userinfo["age"] = 20
	userinfo["hobby"] = []string{"吃饭", "睡觉"}
	// fmt.Println(userinfo) //map[age:20 hobby:[吃饭 睡觉] username:张三]

	// 空接口不能输出 下标
	// fmt.Println(userinfo["hobby"][0]) // userinfo["hobby"][0] (type interface {} does not support indexing)

	var address1 = Address{
		Name:  "李四",
		phone: "13123456789",
	}
	userinfo["address"] = address1
	// fmt.Println(userinfo["address"].Name) //  userinfo["address"].Name undefined (type interface {} is interface with no methods)

	// 通过类型断言打印
	hobby1, _ := userinfo["hobby"].([]string)
	fmt.Println(hobby1[0]) // 吃饭

	address2, _ := userinfo["address"].(Address)
	fmt.Println(address2.Name, address2.phone) // 李四 13123456789
}
