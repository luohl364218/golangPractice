package main

import (
	"net"
	"fmt"
	"golangPractice/chat_room/model"
)
var user model.User

func main() {
	connect, err := net.Dial("tcp", "localhost:10000")
	if err != nil {
		fmt.Println("dial tcp server host err:", err)
		return
	}

	fmt.Println("please input userId and password:")
	fmt.Scanf("%d %s", &user.UserId, &user.Passwd)

	//TCP必须保持连接
	err = login(connect, user.UserId, user.Passwd)
	if err != nil {
		fmt.Println("login failed, err:", err)
		return
	}

	//显示在线用户列表
	showOnlineUserList()
	//并发处理服务端的消息
	go processServerMessage(connect)

	defer connect.Close()

	for {
		//用户交互逻辑
		logic()
	}
}
