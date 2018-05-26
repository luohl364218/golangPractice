package main

import (
	"fmt"
	"golangPractice/chat_room/model"
	"golangPractice/chat_room/protocol"
)

//保存在线好友列表
var onlineUserMap map[int]*model.User = make(map[int]*model.User, 16)

func showOnlineUserList() {
	//显示在线用户列表
	fmt.Println("--------------【online user list】-------------")
	for id, _ := range onlineUserMap{
		fmt.Println("user ", id)
	}
	fmt.Println("-----------------------------------------------")
}

func updateUserStatus(userStatus protocol.UserStatusNotify)  {
	user, ok := onlineUserMap[userStatus.UserId]
	if !ok {
		//不存在 则插入 说明是新上线的用户
		user = &model.User{}
		user.UserId = userStatus.UserId
	}
	user.Status = userStatus.Status
	onlineUserMap[user.UserId] = user
	//判断用户状态
	//如果是下线则通知并删除
	if user.Status == model.UserStatusOffline {
		fmt.Println("【broadcast】user ", user.UserId, " offline")
		delete(onlineUserMap, user.UserId)
	}else {
		fmt.Println("【broadcast】user ", user.UserId, " online")
	}
}