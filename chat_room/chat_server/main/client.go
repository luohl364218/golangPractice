package main

import (
	"net"
	"errors"
	"bytes"
	"encoding/binary"
	"fmt"
	"encoding/json"
	"golangPractice/chat_room/protocol"
)

type Client struct {
	conn net.Conn
	buf [8192]byte
}

func (p *Client) readPackage() (msg *protocol.Message, err error) {
	//首先读取消息长度[0:4]
	n, err := p.conn.Read(p.buf[0:4])
	if n != 4 {
		err = errors.New("read header failed")
		return
	}
	//字节转为整型
	buffer := bytes.NewBuffer(p.buf[0:4])
	var temp int32
	err = binary.Read(buffer, binary.BigEndian, &temp)
	if err != nil {
		fmt.Println("read package len failed")
		return
	}
	//拿到消息长度
	packLen := int(temp)
	//判断消息内容长度是否正确
	n, err = p.conn.Read(p.buf[0:packLen])
	if n != packLen {
		err = errors.New("read body failed")
		return
	}
	//转化为结构体
	err = json.Unmarshal(p.buf[0:packLen], msg)
	if err != nil {
		fmt.Println("unmarshal failed, err:", err)
	}
	return
}

func (p *Client) Process() (err error) {
	for {
		var msg *protocol.Message
		msg, err = p.readPackage()
		if err != nil {
			fmt.Println("read package err:", err)
			return
		}
		err = p.processMsg(msg)
		if err != nil {
			fmt.Println("process msg err:", err)
			return
		}
	}
}

func (p *Client) processMsg(msg *protocol.Message) (err error) {
	switch msg.Cmd {
	case protocol.UserLogin:
		err = p.login(msg)
	case protocol.UserRegister:
		err = p.register(msg)
	default:
		err = errors.New("unsupport message")
		return
	}
	return
}

//登录
func (p *Client) login(msg *protocol.Message) (err error) {
	defer func() {
		p.loginResp(err)
	}()

	var cmd protocol.LoginCmd
	err = json.Unmarshal([]byte(msg.Data), &cmd)
	if err != nil {
		return
	}
	_, err = mgr.Login(cmd.Id, cmd.Passwd)
	if err != nil {
		return
	}
	return
}

//登录回复消息
func (p *Client) loginResp(err error) {
	//【回复】消息结构体
	var respMsg protocol.Message
	//【回复】命令头部
	respMsg.Cmd = protocol.UserLoginRes
	//组装【登录回复】的内容
	var loginRes protocol.LoginCmdRes
	loginRes.Code = protocol.CorrectCode
	if err != nil {
		loginRes.Code = protocol.ErrorCode
		loginRes.Error = fmt.Sprintf("%v", err)
	}
	//序列化【登录回复】消息
	data, err := json.Marshal(loginRes)
	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}
	//【回复】命令内容
	respMsg.Data = string(data)
	//序列化【回复】消息
	data, err = json.Marshal(respMsg)
	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}
	//发送【回复】消息
	err = p.writePackage(data)
	if err != nil {
		fmt.Println("send failed, err:", err)
		return
	}
}

//注册
func (p *Client) register(msg *protocol.Message) (err error) {
	var cmd protocol.RegisterCmd
	err = json.Unmarshal([]byte(msg.Data), &cmd)
	if err != nil {
		return
	}
	err = mgr.Register(&cmd.User)
	if err != nil {
		return
	}
	return
}

//写回复
func (p *Client) writePackage(data []byte) (err error) {
	//写消息长度
	packLen := len(data)
	buffer := bytes.NewBuffer(p.buf[0:4])
	//将长度值写入切片
	err = binary.Write(buffer, binary.BigEndian, packLen)
	if err != nil {
		fmt.Println("write package len failed")
		return
	}
	//发送消息长度
	n, err := p.conn.Write(p.buf[0:4])
	if err != nil {
		err = errors.New("write header failed")
		return
	}
	//将具体消息发送
	n, err = p.conn.Write(data)
	if err != nil {
		fmt.Println("write data failed, err:", err)
		return
	}

	if n != packLen {
		err = errors.New("write data not finished")
		return
	}
	return
}