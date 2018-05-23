package main

import (
	"fmt"
	"time"
)

func main() {

	chann := make(chan int, 1000)
	for i := 0; i < 1000; i++ {
		chann <- i
	}

	//使用select没有阻塞
	for {
		select {
		case v := <-chann:
			fmt.Println(v)
		default:
			fmt.Println("waiting...")
			time.Sleep(time.Second)
		}
	}

}


