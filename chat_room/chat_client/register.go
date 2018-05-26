package main

import (
	"errors"
	"fmt"
	"net"
	"golangPractice/chat_room/protocol"
	"encoding/json"
	"encoding/binary"
)

//注册功能
func register(conn net.Conn, userId int, password string) (err error) {
	var msg protocol.Message
	msg.Cmd = protocol.UserRegister

	var registerCmd protocol.RegisterCmd
	registerCmd.User.UserId = userId
	registerCmd.User.Nick = fmt.Sprintf("stu%d", userId)
	registerCmd.User.Sex = "male"
	registerCmd.User.Passwd = password
	registerCmd.User.Header = fmt.Sprintf("/header/%d.jgp", userId)

	data, err := json.Marshal(registerCmd)
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

	_, err = conn.Write([]byte(data))
	if err != nil {
		return
	}
	//读取服务端回复
	rspMsg, err := readPackage(conn)
	if err != nil {
		fmt.Println("read package failed, err:", err)
	}
	fmt.Println(*rspMsg)
	return
}
