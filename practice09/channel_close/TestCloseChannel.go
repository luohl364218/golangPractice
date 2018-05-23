package main

import "fmt"

func main() {

	chann := make(chan int, 10)
	for i := 0; i < 10; i++{
		chann <- i
	}
	close(chann)

	for {
		v, ok :=  <- chann
		//检测管道是否关闭
		if ok == false {
			fmt.Println("chan is close")
			break
		}
		fmt.Println(v)
	}

}

