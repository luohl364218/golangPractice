package main

import "fmt"

//保存在线好友列表

var onlineUser []int

func showOnlineUserList() {
	//显示在线用户列表
	fmt.Println("online user list:")
	for _, v := range onlineUser{
		fmt.Println("user ", v)
	}
}