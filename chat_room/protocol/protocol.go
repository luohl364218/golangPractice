package protocol

import "golangPractice/chat_room/model"

type Message struct {
	Cmd string `json:"cmd"`
	Data string `json:"data"`
}

type LoginCmd struct {
	Id int `json:"user_id"`
	Passwd string `json:"passwd"`
}

type RegisterCmd struct {
	User model.User `json:"user"`
}

type LoginCmdRes struct {
	Code int `json:"code"`
	Users []int `json:"users"`
	Error string `json:"error"`
}

type UserStatusNotify struct {
	UserId int `json:"user_id"`
	Status int `json:"status"`
}

type UserSendMessageReq struct {
	//发送消息的用户ID
	UserId int `json:"user_id"`
	Data string `json:"data"`
}

type UserRecvMessageReq struct {
	//来自哪个用户
	UserId int `json:"user_id"`
	Data string `json:"data"`
}