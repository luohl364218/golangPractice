package main

import (
	"net"
	"fmt"
	"golangPractice/chat_room/protocol"
	"encoding/json"
	"bytes"
	"encoding/binary"
	"errors"
	"bufio"
	"os"
	"strings"
)

func main() {
	connect, err := net.Dial("tcp", "localhost:10000")
	if err != nil {
		fmt.Println("dial tcp server host err:", err)
		return
	}

	//TCP必须保持连接
	err = login(connect)
	if err != nil {
		fmt.Println("login failed, err:", err)
		return
	}

	defer connect.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		sendStr, _ := reader.ReadString('\n')
		sendStr = strings.Trim(sendStr, "\n\r")
		if sendStr == "Q" {
			return
		}

		_, err := connect.Write([]byte(sendStr))
		if err != nil {
			return
		}
	}
}

func login(conn net.Conn) (err error) {
	var msg protocol.Message
	msg.Cmd = protocol.UserLogin

	var loginCmd protocol.LoginCmd
	loginCmd.Id = 1
	loginCmd.Passwd = "123456789"

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
	buffer := bytes.NewBuffer(buf[:])
	//将长度值写入切片
	err = binary.Write(buffer, binary.BigEndian, packLen)
	if err != nil {
		fmt.Println("write package len failed")
		return
	}
	//发送消息长度
	n, err := conn.Write(buffer.Bytes())
	//todo 64位系统为8 32位系统为4
	if err != nil || n != 8{
		err = errors.New("write header failed")
		return
	}

	_, err = conn.Write([]byte(data))
	if err != nil {
		return
	}
	return
}
