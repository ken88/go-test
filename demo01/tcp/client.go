package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

// NewClient 链接客户端
func NewClient(serverIp string, serverPort int) *Client {
	// 创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
	}

	// 链接服务端
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("服务端链接失败err", err)
		return nil
	}
	client.conn = conn
	// 返回对象
	return client
}

// menu 设置cli模式选择器
func (client *Client) menu() bool {
	var flag int
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")
	fmt.Scanln(&flag)
	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println(">>>>>>请输入合法范围内的数<<<<<<")
		return false
	}
}

// SelectUsers 查询当前在线用户
func (client *Client) SelectUsers() {
	sendMsg := "who\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.write.err", err)
		return
	}
}

// PrivateChat 私聊模式
func (client *Client) PrivateChat() {
	var remoteName string
	var chatMsg string

	// 打印所有用户
	client.SelectUsers()
	fmt.Println(">>>请输入要聊天的对象[用户名]<<<,exit退出")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println(">>>>请输入消息内容<<<<,exit退出")
		fmt.Scanln(&chatMsg)
		for chatMsg != "exit" {
			// 发消息给服务器

			// 消息不为空则发送
			if len(chatMsg) != 0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n"
				_, err := client.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("client.conn.write.err", err)
					break
				}
			}
			chatMsg = ""
			fmt.Println(">>>>请输入消息内容<<<<,exit退出")
			fmt.Scanln(&chatMsg)
		}
		// 打印所有用户
		client.SelectUsers()
		fmt.Println(">>>请输入要聊天的对象[用户名]<<<,exit退出")
		fmt.Scanln(&remoteName)
	}
}

// PublicChat  公聊模式
func (client *Client) PublicChat() {
	// 提示用户输入信息
	var chatMsg string
	fmt.Println(">>>请输入聊天内容<<<,exit退出。")
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		// 发消息给服务器

		// 消息不为空则发送
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("client.conn.write.err", err)
				break
			}
		}
		chatMsg = ""
		fmt.Println(">>>请输入聊天内容<<<,exit退出。")
		fmt.Scanln(&chatMsg)
	}

}

// UpdateName 更改用户名
func (client *Client) UpdateName() bool {
	fmt.Println(">>>>>>请输入用户名<<<<<<<")
	fmt.Scanln(&client.Name)
	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.write.err", err)
		return false
	}
	return true
}

// DealResponse 处理server回应的消息，直接显示标准输出即可
func (client *Client) DealResponse() {
	// 一旦client.conn有数据，就直接copy到stdout标准输出上，永久阻塞监听
	io.Copy(os.Stdout, client.conn)

}

// Run 启动客户端
func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {

		}
		// 根据不同的模式处理不同的业务
		switch client.flag {
		case 1:
			// 公聊模式
			client.PublicChat()
			break
		case 2:
			// 私聊模式
			client.PrivateChat()
			break
		case 3:
			// 修改用户名
			client.UpdateName()
			break

		}
	}
}

var serverIp string
var serverPort int

// ./client -ip 127.0.0.1 -port 8888
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器ip地址(默认是127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口(默认是8888)")
}

func main() {
	// 解析命令行
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>>>>>服务器链接失败....")
		return
	}
	fmt.Println(">>>>>>服务器链接成功.....")

	// 单独开启一个goroutine 去处理server的回应消息
	go client.DealResponse()

	// 启动客户端业务
	client.Run()
}
