package model

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"encoding/json"
	"time"
)

var (
	UserTable = "users"
)

type UserMgr struct {
	pool *redis.Pool
}

func NewUserMgr(pool *redis.Pool) (mgr *UserMgr) {
	mgr = &UserMgr{
		pool:pool,
	}
	return
}

func (p *UserMgr) getUser(conn redis.Conn, id int) (user *User, err error) {
	result, err := redis.String(conn.Do("HGet", UserTable, fmt.Sprintf("%d", id)))
	if err != nil {
		if err == redis.ErrNil {
			err =  ErrUserNotExist
		}
		return
	}

	user = &User{}
	err = json.Unmarshal([]byte(result), user)
	if err != nil {
		return
	}
	return
}

func (p *UserMgr) Login(id int, passwd string) (user *User, err error) {
	//获取Redis连接
	conn := p.pool.Get()
	defer conn.Close()
	//判断是否能从Redis中获取用户对象
	user, err = p.getUser(conn, id)
	if err != nil{
		return
	}
	//判断密码是否正确
	if user.UserId != id || user.Passwd != passwd {
		err = ErrInvalidPassword
	}
	//设置用户登录状态
	user.Status = UserStatusOnline
	user.LastLogin = fmt.Sprintf("%v", time.Now())
	return
}

func (p *UserMgr) Register(user *User) (err error) {
	//判断输入是否为空
	if user == nil {
		err = ErrInvalidParams
		return
	}
	//获取Redis连接
	conn := p.pool.Get()
	defer conn.Close()
	//判断是否能从Redis中获取用户对象
	_, err = p.getUser(conn, user.UserId)
	if err == nil{
		//不报错 说明用户ID已经存在
		err = ErrUerExist
		return
	}else if err != ErrUserNotExist {
		//未知错误，返回
		return
	}
	//json序列化
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	_, err = conn.Do("HSet", UserTable, fmt.Sprintf("%d", user.UserId), string(data))
	if err != nil {
		err = ErrInsertUser
	}
	return
}



