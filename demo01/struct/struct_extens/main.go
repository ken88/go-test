package main

import "fmt"

// 父类
type Animal struct {
	Name string
}

func (an Animal) run() {
	fmt.Printf("%v 在溜达...\n", an.Name)
}

// 子类
type Dog struct {
	Age    int
	Animal Animal // 继承父类
}

func (d Dog) wang() {
	fmt.Printf("%v 在汪汪\n", d.Animal.Name)
}

func main() {
	var a = Dog{
		Age: 20,
		Animal: Animal{
			Name: "小黑",
		},
	}
	a.Animal.run()
	a.wang()
}
