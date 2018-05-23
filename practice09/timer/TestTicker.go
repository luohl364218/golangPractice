package main

import (
	"time"
	"fmt"
)

func main()  {
	//定时器
	v := time.NewTicker(time.Second)

	for {
		select {
		case <-v.C:
			fmt.Println("time out")
		}
	}


}

