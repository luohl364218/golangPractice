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
	_, err = connect.Do("lpush", "book_list", "abc","efg", 100)
	if err != nil {
		fmt.Println("lpush err:", err)
		return
	}

	r, err := redis.String(connect.Do("lpop", "book_list"))
	if err != nil {
		fmt.Println("get book_list err:", err)
		return
	}
	fmt.Println("result:", r)

	r, err = redis.String(connect.Do("lpop", "book_list"))
	if err != nil {
		fmt.Println("get book_list err:", err)
		return
	}
	fmt.Println("result:", r)

	r, err = redis.String(connect.Do("lpop", "book_list"))
	if err != nil {
		fmt.Println("get book_list err:", err)
		return
	}
	fmt.Println("result:", r)
}
