package main

import (
	"fmt"
)

func calc(taskChan , resultChan chan int, exitChan chan bool)  {

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
	exitChan <- true
	fmt.Println("exit calc")
}

func main()  {

	taskChan := make(chan int, 1000)
	resultChan := make(chan int, 1000)
	//使用chan
	exitChan := make(chan bool, 8)

	//新开协程跑匿名函数
	go func() {
		for i := 1; i < 10000; i++ {
			taskChan <- i
		}
		//关闭chan之后range taskChan时会在取完后自动退出
		close(taskChan)
	}()

	//新开8个协程计算
	for i := 0; i < 8; i++ {
		go calc(taskChan, resultChan, exitChan)
	}

	//新开协程取结果
	go func() {
		//等待所有goroute退出
		for i := 0; i < 8; i++ {
			//阻塞住，只要取8个值，说明所有goroute都退出，说明结果已经计算完了
			<- exitChan
			fmt.Println("wait goroute ", i, " exited")
		}
		close(resultChan)
	}()

	for v := range resultChan{
		fmt.Println(v)
	}
}

