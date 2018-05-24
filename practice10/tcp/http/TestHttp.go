package main

import (
	"net"
	"fmt"
	"io"
)

func main() {

	connect, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println("connect err:", err)
		return
	}

	defer connect.Close()

	msg := "Get / HTTP/1.1\r\n"
	msg += "Host:www.baidu.com\r\n"
	msg += "Connection:close\r\n"
	msg += "\r\n\r\n"

	_, err = io.WriteString(connect, msg)
	if err != nil {
		fmt.Println("write string failed. ",err)
		return
	}
	buf := make([]byte, 4096)
	for {
		count, err := connect.Read(buf)
		if err != nil {
			fmt.Println("read err:", err)
			break
		}
		fmt.Println(string(buf[:count]))
	}
}


