package main

import (
	"time"
	"fmt"
)

func main() {
	//定时一次
	select {
	case <-time.After(time.Second):
		fmt.Println("after second")
	}
}