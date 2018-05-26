package main

import (
	"golangPractice/chat_room/protocol"
	"fmt"
	"net"
	"encoding/json"
	"encoding/binary"
	"errors"
	"os"
	"golangPractice/chat_room/model"
)

//注意 先在Redis中执行 hset users 1 "{\"user_id\":1,\"passwd\":\"123456789\"}" 存储用户信息
func login(conn net.Conn, userId int, password string) (err error) {
	var msg protocol.Message
	msg.Cmd = protocol.UserLogin

	var loginCmd protocol.LoginCmd
	loginCmd.Id = userId
	loginCmd.Passwd = password

	data, err := json.Marshal(loginCmd)
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
	fmt.Println("recv login resp msg:", *rspMsg)
	//转为返回结果
	var loginResp protocol.LoginCmdRes
	err = json.Unmarshal([]byte(rspMsg.Data), &loginResp)
	if err != nil {
		fmt.Println("unmarshal login response err:", err)
		return
	}
	//如果返回错误，执行注册
	if loginResp.Code == protocol.ErrorCode{
		fmt.Println("user not register, start register")
		register(conn, userId, password)
		os.Exit(0)
	}
	//保存在线用户列表
	for _, v := range loginResp.Users{
		if v == userId {
			continue
		}
		//本地内存中保存
		user := &model.User{
			UserId:v,
		}
		onlineUserMap[v] = user
	}
	return
}

