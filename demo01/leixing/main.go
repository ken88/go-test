package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	//int 8
	// var num int8 = 130 // constant 130 overflows int8 （超出 int8 -127 到 127）
	// fmt.Printf("num=%v 类型%T", num, num)

	//unsafe.Sizeof() 查看不同长度的整型 在内存里面的存储空间
	// var a int8 = 15
	// fmt.Printf("num=%v 类型%T \n", a, a)
	// fmt.Println(unsafe.Sizeof(a)) // 1字节

	// int 类型转换
	// var a1 int16 = 12
	// var a2 int32 = 29
	// fmt.Println(a1 + int16(a2))

	// 保留3位小数输出
	// var c float64 = 3.123456
	// fmt.Printf("%.3f", c)

	// golang中float 精度丢失问题
	f1 := 1219.1
	fmt.Println("1219.1 x 100 =", f1*100) // 输出值121909.99999999999 正常是121910
	d1 := decimal.NewFromFloat(f1).Mul(decimal.NewFromInt(100))
	fmt.Println("1219.1 x 100 =", d1) // 121910

	m1 := 8.2
	m2 := 3.8
	fmt.Println("8.3 - 3.8 =", m1-m2) // 输出值4.3999999999999995 正常值 4.4
	d2 := decimal.NewFromFloat(m1).Sub(decimal.NewFromFloat(m2))
	fmt.Println("8.3 - 3.8 =", d2) // 4.4

	m3 := 8.1
	m4 := 3.8
	fmt.Println("8.1 + 3.8 =", m3+m4) // 输出值11.899999999999999 正常值 11.9
	d3 := decimal.NewFromFloat(m3).Add(decimal.NewFromFloat(m4))
	fmt.Println("8.1 + 3.8 =", d3) // 11.9

}
