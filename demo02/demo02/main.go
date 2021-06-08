package main

func main()  {
	// 面这段代码输出什么，说明原因。
	slice := []int{0,1,2,3}
	m:= make(map[int]*int)
	for key,val :=range slice{
		m[k]=&val
	}
}