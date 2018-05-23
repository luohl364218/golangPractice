package main

import (
	"fmt"
	"time"
)

func calc(taskChan , resultChan chan int)  {

	//判断是不是素数
	for v := range taskChan{
		flag := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			resultChan <- v
		}

	}
}

func main()  {

	taskChan := make(chan int, 1000)
	resultChan := make(chan int, 1000)

	//新开协程跑匿名函数
	go func() {
		for i := 1; i < 10000; i++ {
			taskChan <- i
		}
		close(taskChan)
	}()

	//新开8个协程计算
	for i := 0; i < 8; i++ {
		go calc(taskChan, resultChan)
	}

	//新开协程取结果
	go func() {
		for {
			v := <- resultChan
			fmt.Println(v)
		}
	}()

	time.Sleep(2 * time.Second)
}

