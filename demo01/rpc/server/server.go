package main

import (
	"gotest/demo01/rpc/param" // 自定义包结构体
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
}

// 计算加法
func (t *Args) Add(param param.Parram, reply *int) error {
	*reply = param.X + param.Y // 加法计算
	return nil                 // 返回类型
}

// 计算乘法
func (t *Args) Cheng(param param.Parram, reply *int) error {
	*reply = param.X * param.Y // 乘法计算
	return nil                 // 返回类型
}

func main() {
	// 1. 初始化指针数据类型
	arith := new(Args)

	// 2. 掉用net/rpc 包的功能 将服务对象进行注册
	err := rpc.Register(arith)
	if err != nil {
		panic(err.Error())
	}
	// 3. 通过该函数把Args中提供的服务注册到http请求上，方便调用者利用http方式进行数据传递
	rpc.HandleHTTP()

	// 4. 在特定的端口进行监听
	l, e := net.Listen("tcp", "127.0.0.1:1234")
	if e != nil {
		panic(e.Error())
	}

	// 5. 开启服务
	http.Serve(l, nil)
	// go http.Serve(l, nil) // 后台运行
}
