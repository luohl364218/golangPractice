package main

import (
	"net/http"
	"fmt"
)

var urlStr = []string{
	"http://www.baidu.com",
	"http://www.google.com",
	"http://www.taobao.com",
}

func main() {

	for _, v := range urlStr {
		resp, err := http.Head(v)
		if err != nil {
			fmt.Printf("head %s failed,err:%v\n", v, err)
			continue
		}
		fmt.Printf("head success, status: %v\n", resp.Status)
	}

}
