package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	// 在线当前用户列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播的channel
	Message chan string
}

// NewServer 创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// Start 启动服务器的接口
func (this *Server) Start() {

	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("链接错误", err)
	}

	// close socket
	defer listener.Close()

	// 启动监听Message的goroutine
	go this.ListenMessage()

	for {
		// 循环接入所有客户端得到专线连接 Accept
		conn, e := listener.Accept()
		if e != nil {
			fmt.Println("listener accept error ", err)
			continue
		}

		// 开辟独立协程与该客聊天
		go this.Handler(conn)
	}
}

// Handler 接收用户的消息
func (this *Server) Handler(conn net.Conn) {

	user := NewUser(conn, this)

	// 用户登陆，将用户加入到OnlineMap中
	user.Online()

	// 监听用户是否是活跃的channel
	isLive := make(chan bool)

	// 接收客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				// 用户下线
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("conn read err:", err)
				return
			}

			// 提取用户的消息（去除'\n'）
			msg := string(buf[:n-1])

			// 用户针对msg进行消息处理
			user.Domessage(msg)

			// 用户的任意消息，代表当前用户是一个活跃的
			isLive <- true
		}
	}()

	for {
		// 当前handler阻塞
		select {
		case <-isLive:
		// 当前用户是活跃的，应该重置定时器
		// 不做任何事情
		case <-time.After(time.Second * 300):
			// 已经超时
			// 将当前的User强制关闭
			user.SendMsg("你被踢了\n")

			// 销毁用户的资源
			close(user.C)

			// 关闭链接
			conn.Close()

			// 退出当前Handler
			return
		}
	}

}

// BroadCast 广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	senMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- senMsg
}

// ListenMessage 监听Message广播消息channel的goroutine,一旦有消息就发送给当前在线的用户
func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message

		// 将msg发送给全部的在线user
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}
