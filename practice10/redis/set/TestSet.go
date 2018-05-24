package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func main() {
	//DoSet()
	DoManySet()

}

func DoSet()  {
	connect, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("connect redis err:", err)
		return
	}
	defer connect.Close()
	_, err = connect.Do("Set", "abc", 100)
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
}

//批量set
func DoManySet()  {
	connect, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("connect redis err:", err)
		return
	}
	defer connect.Close()
	_, err = connect.Do("MSet", "abc", 100, "efg", 200)
	if err != nil {
		fmt.Println("sets err:", err)
		return
	}

	r, err := redis.Ints(connect.Do("MGet", "abc", "efg"))
	if err != nil {
		fmt.Println("get abc err:", err)
		return
	}

	for _,v := range r{
		fmt.Println("result:", v)
	}
}