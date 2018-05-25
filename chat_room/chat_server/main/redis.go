package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
	"fmt"
)

var pool *redis.Pool

func initRedis(addr string, maxIdle int, maxActive int, idleTimeout time.Duration) {

	pool = &redis.Pool{
		MaxIdle:maxIdle,
		MaxActive:maxActive,
		IdleTimeout:idleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
	}
	fmt.Println("init redis success")
	return
}

func GetConn() redis.Conn {
	return pool.Get()
}

func PutConn(conn redis.Conn) {
	conn.Close()
}
