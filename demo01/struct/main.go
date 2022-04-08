package main

import "fmt"

func main() {
	// 定义结构体
	// f1()

	// 结构体是值类型 验证
	// checkType()
}

// 结构体是值类型
func checkType() {
	type persion struct {
		username string
		age      int
	}
	var p1 persion
	p1.username = "张三"
	p1.age = 20
	fmt.Println(p1) // {张三 20}

	p2 := p1
	p2.username = "李四"
	fmt.Println(p2) // {李四 20}
	fmt.Println(p1) // {张三 20}

}

// 定义结构体
func f1() {
	// 1 定义结构体
	type persion struct {
		username string
		age      int
		sex      string
	}

	// 2.赋值 ：扑通结构体类型 类型main.Persion
	// var p1 persion
	// p1.username = "张三"
	// p1.age = 20
	// p1.sex = "男"
	// fmt.Printf("值%v 类型%T\n", p1, p1)  // 值{张三 20 男} 类型main.Persion
	// fmt.Printf("值%#v 类型%T\n", p1, p1) // 值main.Persion{username:"张三", age:20, sex:"男"} 类型main.Persion

	//3.第二种赋值 ：结构体指针 类型*main.persion
	// var p2 = new(persion)
	// p2.username = "张三"
	// p2.age = 20
	// p2.sex = "男"
	// (*p2).username = "lisi"           // p2.username p2.age p2.sex 这三种指针变量赋值后，最后也会改成 (*p2).username 这种方式改变
	// fmt.Printf("值%v 类型%T\n", p2, p2)  // 值&{lisi 20 男} 类型*main.persion
	// fmt.Printf("值%#v 类型%T\n", p2, p2) // 值&main.persion{username:"lisi", age:20, sex:"男"} 类型*main.persion

	//3.第三种赋值 ：结构体指针 类型*main.persion
	// var p3 = &persion{}
	// p3.username = "小七"
	// p3.age = 25
	// p3.sex = "女"
	// fmt.Printf("值%#v 类型%T\n", p3, p3) // 值&main.persion{username:"小七", age:25, sex:"女"} 类型*main.persion

	//4 .第四种赋值 ：扑通结构体类型 类型main.persion
	// p4 := persion{
	// 	username: "赵六",
	// 	age:      44,
	// 	sex:      "男",
	// }
	// fmt.Printf("值%#v 类型%T\n", p4, p4) // 值main.persion{username:"赵六", age:44, sex:"男"} 类型main.persion

	//5 .第五种赋值 结构体指针 类型*main.persion
	// p5 := &persion{
	// 	username: "赵六",
	// 	age:      44,
	// 	sex:      "男",
	// }
	// fmt.Printf("值%#v 类型%T\n", p5, p5) // 值&main.persion{username:"赵六", age:44, sex:"男"} 类型*main.persion
}
