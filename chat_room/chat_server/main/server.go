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
	client := &Client{
		conn:conn,
	}
	err := client.Process()
	if err != nil {
		fmt.Println("client process err, ", err)
		return
	}
}
