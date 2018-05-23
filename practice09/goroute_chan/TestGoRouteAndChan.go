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

	fmt.Println("----------------------")

	test()
}

func test() {
	ch := make(chan int,10)
	exitCh := make(chan struct{}, 2)
	go send(ch, exitCh)
	go recv(ch, exitCh)

	//等待所有goroute退出【采用计数方法比较稳】
	for i := 0; i < 2; i++{
		<- exitCh
	}
}

func send(ch chan int, exitCh chan struct{})  {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("generate ", i)
	}
	//记得关闭
	close(ch)
	exitCh <- struct{}{}
}

func recv(ch chan int, exitCh chan struct{}) {
	for {
		v, ok := <- ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	exitCh <- struct{}{}
}