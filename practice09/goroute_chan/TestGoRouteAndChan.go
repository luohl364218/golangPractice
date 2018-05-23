package main

import (
	"fmt"
	"time"
)

func write(ch chan int)  {
	for i := 0; i < 100; i++ {
		ch <- i
		fmt.Println("generate ", i)
	}
}

func read(ch chan int)  {
	for {
		b := <-ch
		fmt.Println(b)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	chann := make(chan int, 10)
	go write(chann)
	go read(chann)
	time.Sleep(5 * time.Second)
}