package main

import "fmt"

type Persion struct {
	Username string
	Age      int
	Sex      string
	Height   float64
}

// 定义结构体方法
func (p Persion) printInfo() {
	fmt.Printf("姓名：%v 年龄：%v 性别：%v 身高：%v\n", p.Username, p.Age, p.Sex, p.Height)
}

// 修改属性值
func (p *Persion) setPersion(username string, age int) {
	p.Username = username
	p.Age = age
}

func main() {

	// 1.
	// var p1 = Persion{
	// 	Username: "张三",
	// 	Age:      20,
	// 	Sex:      "男",
	// 	Height:   1.76,
	// }
	// p1.printInfo() // 姓名：张三 年龄：20 性别：男 身高：1.76

	// var p2 = Persion{
	// 	Username: "李四",
	// 	Age:      25,
	// 	Sex:      "女",
	// 	Height:   1.66,
	// }
	// p2.printInfo() // 姓名：李四 年龄：25 性别：女 身高：1.66

	//2.
	var p1 = Persion{
		Username: "张三",
		Age:      20,
		Sex:      "男",
		Height:   1.76,
	}
	p1.printInfo() // 姓名：张三 年龄：20 性别：男 身高：1.76
	// 修改属性值
	p1.setPersion("李四", 30)
	p1.printInfo() // 姓名：李四 年龄：30 性别：男 身高：1.76
}
