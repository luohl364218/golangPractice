package main

import (
	"net"
	"fmt"
)

func runServer(addr string) (err error) {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
			continue
		}
		go handleConnect(conn)
	}
}

func handleConnect(conn net.Conn)  {
	defer conn.Close()
	for {
		buf := make([]byte, 512)
		//记录读取的数量，不然会输出很多空白的内容
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err:", err)
			return
		}
		fmt.Printf(string(buf[:n]))
		//todo
	}
}
