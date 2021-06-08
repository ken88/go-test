package main

import (
	"fmt"
	"sort"
)

func main() {
	//通过mack创建切片
	//makeSlice()

	// 检测切片是否是引用类型
	// checkType()

}

// 检测切片是否是引用类型  潜copy 复制的是地址
func checkType() {
	/*
		创建切片的时候 会生成一个数组，例如数组的 地址为 0x...010

		s1保存的是数组的地址 0x...010
	*/
	s1 := []int{1, 2, 3, 4}

	fmt.Println(s1) //[1 2 3 4]

	/*
		s2 := s1
		s2 := 0x...010
		将s1保存的地址赋值给s2  当前s1与s2保存的都是 内存地址 0x...010
	*/
	s2 := s1 //将s1的地址赋值给s2

	fmt.Println(s2) //[1 2 3 4]
	s2[1] = 22
	fmt.Println(s1) //[1 22 3 4]  s1下标1的值也更改了22 所以 切片是引用类型
	fmt.Println(s2) //[1 22 3 4]
}

// 切片 - 排序
func qiepian() {

	var numSlice = []int{9, 28, 47, 76, 45, 34, 33} //定义切片
	fmt.Println(numSlice)
	for i := 0; i < len(numSlice); i++ {
		for j := i + 1; j < len(numSlice); j++ {

			if numSlice[i] > numSlice[j] {
				tmp := numSlice[i]
				numSlice[i] = numSlice[j]
				numSlice[j] = tmp
			}
		}

	}
	sort.Ints(numSlice) //升序
	fmt.Println(numSlice)
	sort.Sort(sort.Reverse(sort.IntSlice(numSlice))) //降序
	fmt.Println(numSlice)
}

// 通过mack创建切片
func makeSlice() {
	//定义make类型 int型切片
	a := make([]int, 4, 8)
	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	fmt.Println(a)                                      //[1 2 3 4]
	fmt.Printf("长度%v 容量%v 地址：%p\n", len(a), cap(a), &a) //长度4 容量8 地址：0xc0000044a0

	a = append(a, 5)
	fmt.Printf("长度%v 容量%v 地址：%p\n", len(a), cap(a), &a) //长度5 容量8 地址：0xc0000044a0

	a = append(a, 6)
	fmt.Printf("长度%v 容量%v 地址：%p\n", len(a), cap(a), &a) //长度6 容量8 地址：0xc0000044a0

	a = append(a, 7)
	fmt.Printf("长度%v 容量%v 地址：%p\n", len(a), cap(a), &a) //长度7 容量8 地址：0xc0000044a0

	a = append(a, 8)
	fmt.Printf("长度%v 容量%v 地址：%p\n", len(a), cap(a), &a) //长度8 容量8 地址：0xc0000044a0

	a = append(a, 9)
	fmt.Printf("长度%v 容量%v 地址：%p\n", len(a), cap(a), &a) //长度9 容量16 地址：0xc0000044a0

	a = append(a, 10, 11, 12, 13, 14, 15, 16, 17)
	fmt.Printf("长度%v 容量%v 地址：%p\n", len(a), cap(a), &a) //长度17 容量32 地址：0xc0000044a0
}
