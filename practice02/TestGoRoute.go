package practice02

import (
	"fmt"
	"time"
)
//并发goroute
func TestGoRoute() {
	go add(300, 300)
	//需要睡眠，不然主进程退出，看不到go起线程的效果
	time.Sleep(time.Second)
	for i := 0; i < 100; i++ {
		go TestPrint(i)
	}
	time.Sleep(time.Second)
}

func add(a int, b int) {
	sum := a + b
	fmt.Printf("sum = %d",sum)
}

func TestPrint(i int)  {
	fmt.Println(i)
}
