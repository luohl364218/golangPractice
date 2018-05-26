package main

import (
	"net"
	"os"
	"fmt"
	"golangPractice/chat_room/protocol"
	"encoding/json"
)

func processServerMessage(conn net.Conn) {
	//死循环，一直接收
	for {
		msg, err := readPackage(conn)
		if err != nil {
			fmt.Println("read msg err:", err)
			os.Exit(0)
		}

		switch msg.Cmd {
		case protocol.UserStatusNotifyRes:
			var userStatus protocol.UserStatusNotify
			err = json.Unmarshal([]byte(msg.Data), &userStatus)
			if err != nil {
				fmt.Println("unmarshal failed, err:", err)
				return
			}
			updateUserStatus(userStatus)
		case protocol.UserRecvMessageCmd:
			recvMessageFromServer(msg)
		}
	}
}

func recvMessageFromServer(msg *protocol.Message)  {
	var recvMsg protocol.UserRecvMessageReq
	err := json.Unmarshal([]byte(msg.Data), &recvMsg)
	if err != nil {
		fmt.Println("unmarshal failed, err:", err)
		return
	}
	//1.直接显示 可能会有多线程问题
	fmt.Println("【broadcast】user ",recvMsg.UserId," say:", recvMsg.Data)

	//2.后台线程 通过管道将消息给前台
	msgChan <- recvMsg
}