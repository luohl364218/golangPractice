package main

import (
	"time"
	"fmt"
)

func main()  {
	//日期格式化输出
	now := time.Now()
	//必须得用这个字符串格式化
	fmt.Println(now.Format("2006/01/02 15:04:05"))

	//统计程序耗时
	start := time.Now().UnixNano()
	test()
	end := time.Now().UnixNano()

	fmt.Println("costTime:", (end - start) / 1000000)
}

func test()  {
	time.Sleep(time.Millisecond * 100)
}
