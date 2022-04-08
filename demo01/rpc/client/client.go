package main

import (
	"fmt"
	"gotest/demo01/rpc/param"
	"log"
	"net/rpc"
)

func main() {
	// rpc 加法
	// rpcAdd()

	// rpc 乘法
	rpcCheng()
}

// rpc 乘法
func rpcCheng() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err.Error())
	}
	// Synchronous call
	args := &param.Parram{X: 7, Y: 8}
	var reply int

	//  掉用rpc 乘法
	err = client.Call("Args.Cheng", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Println(reply)
}

// rpc 加法
func rpcAdd() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err.Error())
	}
	// Synchronous call
	args := &param.Parram{X: 7, Y: 8}
	var reply int

	// 掉用rpc 加法
	err = client.Call("Args.Add", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Println(reply)
}
