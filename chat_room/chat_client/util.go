package main

import (
	"net"
	"golangPractice/chat_room/protocol"
	"errors"
	"encoding/binary"
	"encoding/json"
	"fmt"
)

//与服务端读写
func readPackage(conn net.Conn) (msg *protocol.Message, err error) {
	var buf [8192]byte
	n, err := conn.Read(buf[0:4])
	if n != 4 {
		err = errors.New("read header failed")
		return
	}
	//字节转为整型
	var temp uint32
	temp = binary.BigEndian.Uint32(buf[0:4])
	//拿到消息长度
	packLen := int(temp)
	//判断消息内容长度是否正确
	n, err = conn.Read(buf[0:packLen])
	if n != packLen {
		err = errors.New("read body failed")
		return
	}
	//转化为结构体
	msg = new(protocol.Message)
	err = json.Unmarshal(buf[0:packLen], msg)
	if err != nil {
		fmt.Println("unmarshal failed, err:", err)
	}
	return
}