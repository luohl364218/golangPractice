package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

//过期时间
func main() {

	connect, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("connect redis err:", err)
		return
	}
	defer connect.Close()
	_, err = connect.Do("expire", "abc", 100)
	if err != nil {
		fmt.Println("expire err:", err)
		return
	}
}

