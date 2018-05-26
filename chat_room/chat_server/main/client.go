package main

import (
	"net"
	"errors"
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
	//todo 64位系统为8 32位系统为4
	//首先读取消息长度[0:8]
	n, err := p.conn.Read(p.buf[0:4])
	if n != 4 {
		err = errors.New("read header failed")
		return
	}
	//字节转为整型
	var temp uint32
	temp = binary.BigEndian.Uint32(p.buf[0:4])
	//拿到消息长度
	packLen := int(temp)
	//判断消息内容长度是否正确
	n, err = p.conn.Read(p.buf[0:packLen])
	if n != packLen {
		err = errors.New("read body failed")
		return
	}
	//转化为结构体
	msg = new(protocol.Message)
	err = json.Unmarshal(p.buf[0:packLen], msg)
	if err != nil {
		fmt.Println("unmarshal failed, err:", err)
	}
	return
}

func (p *Client) Process() (err error) {
	//该死循环不能随便退出，不然后面就没办法与客户端继续通信了
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
			continue
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

	fmt.Printf("recv user login request, data:%v\n", *msg)
	var cmd protocol.LoginCmd
	err = json.Unmarshal([]byte(msg.Data), &cmd)
	if err != nil {
		return
	}
	user, err := mgr.Login(cmd.Id, cmd.Passwd)
	if err != nil {
		return
	}
	fmt.Println("user ", user.UserId, " login success")
	clientMgr.AddClient(cmd.Id, p)
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
		fmt.Println("register user err:", err)
		return
	}
	fmt.Println("user", cmd.User.UserId, " register success")
	return
}

//写回复
func (p *Client) writePackage(data []byte) (err error) {
	//写消息长度
	packLen := uint32(len(data))
	//将长度值写入切片
	binary.BigEndian.PutUint32(p.buf[0:4], packLen)
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

	if n != int(packLen) {
		err = errors.New("write data not finished")
		return
	}
	return
}