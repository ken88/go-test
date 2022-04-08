package main

import "fmt"

func main() {

	//二维数组
	a1 := [...][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(a1)       //[[1 2 3 4] [5 6 7 8] [9 10 11 12]]
	fmt.Println(a1[0])    //[1 2 3 4]
	fmt.Println(a1[0][3]) //4

	// a2 := [...]int{11, 22, 33}
	// a3 := a2
	// a2[0] = 111
	// fmt.Println(a2)
	// fmt.Println(a3)
	// shuzu()
}

//一维数组
func arr1() {
	var arr1 [4]int
	arr1[0] = 1
	arr1[1] = 2
	fmt.Printf("数组的长度：%v 容量：%v \n", len(arr1), cap(arr1)) //数组的长度：4 容量：4
	fmt.Printf("内存地址：%p", &arr1)                          //内存地址：0xc000010400[0 0 0 0]

	var arr2 [4]int
	fmt.Println(arr2) //[0 0 0 0]

	var arr3 = [4]int{1, 2, 3, 4}
	fmt.Println(arr3) //[1 2 3 4]

	var arr4 = [5]int{1, 2, 3}
	fmt.Println(arr4) // [1 2 3 0 0]

	//数组下标1=>3, 3=>2 其余下标值为0
	var arr5 = [5]int{1: 1, 3: 2}
	fmt.Println(arr5) // [0 1 0 2 0]

	// 值不全补空“”
	var arr6 = [5]string{"a", "b", "c"}
	fmt.Println(arr6) // [a b c  ]

	// 可改变数组
	var arr7 = [...]int{1, 2, 3}
	fmt.Println(arr7) //[1 2 3]
}

// 数组
func shuzu() {
	var str1 = [...]int{1, 2, 3, 4}

	for _, val := range str1 {
		fmt.Println(val)
	}

	var str2 = [3]int{1, 2, 3}
	for i := 0; i < len(str2); i++ {
		fmt.Println(str2[i])
	}

	var str3 [4]int
	str3[0] = 1
	str3[1] = 2
	str3[2] = 3
	str3[3] = 4

	for i := 0; i < len(str3); i++ {
		fmt.Println(str3[i])
	}
}
