package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func main() {

	connect, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("connect redis err:", err)
		return
	}
	defer connect.Close()
	_, err = connect.Do("HSet", "books", "abc", 100)
	if err != nil {
		fmt.Println("set err:", err)
		return
	}

	r, err := redis.Int(connect.Do("HGet", "books","abc"))
	if err != nil {
		fmt.Println("get abc err:", err)
		return
	}
	fmt.Println("result:", r)
}



