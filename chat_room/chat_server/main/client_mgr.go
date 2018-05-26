package main

import "fmt"

//所有在线客户端列表
type ClientMgr struct {
	onlineUsers map[int]*Client
}

var(
	clientMgr *ClientMgr
)

func init()  {
	clientMgr = &ClientMgr{
		onlineUsers:make(map[int]*Client, 1024),
	}
}

func (p *ClientMgr) AddClient(userId int, client *Client)  {
	p.onlineUsers[userId] = client
}

func (p *ClientMgr) GetClient(userId int) (client *Client, err error) {
	client, ok := p.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("user %d not exist", userId)
	}
	return
}