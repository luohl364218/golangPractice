package main

import (
	"fmt"
	"net"
	"golangPractice/chat_room/protocol"
	"encoding/json"
	"encoding/binary"
	"errors"
	"os"
)

//客户端交互的逻辑都在这里
func logic(conn net.Conn)  {
	enterMenu(conn)
}

func enterMenu(conn net.Conn)  {
	fmt.Println("----------【please make your choice】----------")
	fmt.Println("1.list online user")
	fmt.Println("2.talk")
	fmt.Println("3.show unread msg")
	fmt.Println("4.exit")
	fmt.Println("-----------------------------------------------")

	var sel int
	fmt.Scanf("%d", &sel)
	switch sel {
	case 1:
		showOnlineUserList()
	case 2:
		enterTalk(conn)
	case 3:
		showUnreadMsg()
	case 4:
		os.Exit(0)
		return
	}
}

func enterTalk(conn net.Conn) {
	var msg string
	fmt.Println("-----------【please input test】-------------")
	fmt.Scanf("%s", &msg)
	sendTextMsg(conn,msg)
}
//发送广播消息
func sendTextMsg(conn net.Conn, text string)  {
	var msg protocol.Message
	msg.Cmd = protocol.UserSendMessageCmd

	var sendReq protocol.UserSendMessageReq
	sendReq.Data = text
	sendReq.UserId = user.UserId

	data, err := json.Marshal(sendReq)
	if err != nil {
		return
	}
	msg.Data = string(data)
	data, err = json.Marshal(msg)
	if err != nil {
		return
	}
	//写消息长度
	var buf [4]byte
	packLen := uint32(len(data))
	//将长度值写入切片
	binary.BigEndian.PutUint32(buf[:], packLen)
	//发送消息长度
	n, err := conn.Write(buf[:])
	if err != nil || n != 4{
		err = errors.New("write header failed")
		return
	}
	//数据发送
	_, err = conn.Write([]byte(data))
	if err != nil {
		return
	}
}

func showUnreadMsg() {
	select {
		case msg := <-msgChan:
		fmt.Println("【broadcast】user ",msg.UserId," say:", msg.Data)
	default:
		return
	}
}