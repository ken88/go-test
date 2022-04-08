package main

import (
	"fmt"
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	server *Server
}

// NewUser 创建一个用户的api
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String() // 获取客户端的ip
	fmt.Println("用户信息", userAddr)
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}

	// 启动监听当前用户user channel消息的goroutine
	go user.ListenMessage()
	return user
}

// ListenMessage 监听当前User channel的方法，一旦有消息，就直接发送给客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		// 接收其它用户发送的消息
		this.conn.Write([]byte(msg + "\n"))
	}
}

// Online 用户的上线业务
func (this *User) Online() {
	// 当前用户，将用户加入到OnlineMap中
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	// 广播当前用户上线消息
	this.server.BroadCast(this, "已上线")
}

// Offline 用户下限的业务
func (this *User) Offline() {
	// 用户下线，将用户从OnlineMap中删除
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	// 广播当前用户上线消息
	this.server.BroadCast(this, "已下线")
}

// SendMsg 向当前用户发送消息
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// Domessage 用户处理消息的业务
func (this *User) Domessage(msg string) {
	// 如果客户端发送who，查询当前在线的用户
	if msg == "who" {
		// 查询当前在线的用户有哪些
		for _, v := range this.server.OnlineMap {
			onlineMsg := "[" + v.Name + "]在线...\n"
			this.SendMsg(onlineMsg)
		}
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 消息格式: rename|zhangsan
		newName := strings.Split(msg, "|")[1]
		// 判断当前用户名是否存在
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMsg("当前用户已被使用")
		} else {
			// 更新用户名
			this.server.mapLock.Lock()
			// 删除当前用户名
			delete(this.server.OnlineMap, this.Name)
			// 更改当前用户名
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()
			this.Name = newName
			this.SendMsg("您的用户名已经更改为:" + this.Name + "\n")
		}

	} else if len(msg) > 4 && msg[:3] == "to|" {
		// 消息内: to|张三|你好啊

		// 1. 获取对方用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.SendMsg("消息格式不正确,请使用 to|张三|消息内容 格式 \n")
			return
		}
		// 2. 根据用户名，得到对方User对象
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg("用户名不存在\n")
			return
		}
		// 3. 获取消息内容，通过对方的user对象 将消息内容发送过去
		content := strings.Split(msg, "|")[2]
		if content == "" {
			this.SendMsg("无消息内容\n")
			return
		}
		remoteUser.SendMsg(this.Name + "对你说:" + content + "\n")
	} else {
		this.server.BroadCast(this, msg)
	}

}
