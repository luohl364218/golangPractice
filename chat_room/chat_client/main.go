package main

import (
	"net"
	"fmt"
	"golangPractice/chat_room/protocol"
	"encoding/json"
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

	var userId int
	var password string
	fmt.Println("please input userId and password:")
	fmt.Scanf("%d %s", &userId, &password)

	//TCP必须保持连接
	err = login(connect, userId, password)
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
	}
	return
}

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