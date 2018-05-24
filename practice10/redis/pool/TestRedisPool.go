package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:16,
		MaxActive:1024,
		IdleTimeout:300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {

	connect := pool.Get()
	defer connect.Close()

	_, err := connect.Do("Set", "abc", 100)
	if err != nil {
		fmt.Println("set err:", err)
		return
	}

	r, err := redis.Int(connect.Do("Get", "abc"))
	if err != nil {
		fmt.Println("get abc err:", err)
		return
	}
	fmt.Println("result:", r)
	pool.Close()
}


