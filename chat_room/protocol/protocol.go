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
