package main

import "fmt"

func main() {
	fmt.Println("我是for range输出")
	var arr = []string{"php", "java", "golang", "中国"}
	for _, v := range arr {
		fmt.Println(v)
	}
	fmt.Println("")
	fmt.Println("我是for输出")
	count := len(arr)
	for i := 0; i < count; i++ {
		fmt.Println(arr[i])
	}

	// 跳出多层循环
xxxx:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for n := 0; n < 2; n++ {
				if j == 2 {
					break xxxx
				}
				fmt.Println("i=", i, "j=", j, "n=", n)
			}
		}
	}
}
