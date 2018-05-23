package main

import (
	"time"
	"fmt"
)

func main()  {
	//定时器
	/*v := time.NewTicker(time.Second)

	for {
		select {
		case <-v.C:
			fmt.Println("time out")
		}
	}*/

	//死循环
	t := time.NewTicker(time.Second)
	for v := range t.C {
		fmt.Println("hello, ", v)
	}
	//退出时需要stop 防止内存泄漏
	t.Stop()
}

