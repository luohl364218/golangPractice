package main

import (
	"net/http"
	"fmt"
	"time"
)

var urlStr = []string{
	"http://www.baidu.com",
	"http://www.google.com",
	"http://www.taobao.com",
}

func main() {

	//test()
	testClient()

}

func test() {
	for _, v := range urlStr {
		resp, err := http.Head(v)
		if err != nil {
			fmt.Printf("head %s failed,err:%v\n", v, err)
			continue
		}
		fmt.Printf("head success, status: %v\n", resp.Status)
	}
}

func testClient() {
	for _, v := range urlStr {
		//自定义客户端结构体
		c := http.Client{
			Timeout:time.Second * 2,
		}

		resp, err := c.Head(v)
		if err != nil {
			fmt.Printf("head %s failed,err:%v\n", v, err)
			continue
		}
		fmt.Printf("head success, status: %v\n", resp.Status)
	}
}
