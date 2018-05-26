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

		var userStatus protocol.UserStatusNotify
		err = json.Unmarshal([]byte(msg.Data), &userStatus)
		if err != nil {
			fmt.Println("unmarshal failed, err:", err)
			return
		}

		switch msg.Cmd {
		case protocol.UserStatusNotifyRes:
			updateUserStatus(userStatus)
			showOnlineUserList()
		}
	}
}