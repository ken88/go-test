package main

import (
	"fmt"
	"runtime"
)

func init() {
	// 获取逻辑cpu的数量
	fmt.Println("逻辑CPU的书里数量->", runtime.NumCPU()) // 逻辑CPU的书里数量-> 8

	// 设置go程序执行的最大cpu数量：【1，256】
	n := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(n)
}

func main() {

	// 获取goroot目录
	fmt.Println("GOROOT->", runtime.GOROOT()) // GOROOT-> /usr/local/go

	// 获取操作系统
	fmt.Println("os/platform->", runtime.GOOS) // os/platform darwin

}
