package main

import "fmt"

func main() {

	chann := make(chan int, 1000)
	for i := 0; i < 1000; i++ {
		chann <- i
	}
	//关闭管道
	close(chann)
	//range遍历时如果管道关闭会主动停止
	for v := range chann{
		fmt.Println(v)
	}

}
