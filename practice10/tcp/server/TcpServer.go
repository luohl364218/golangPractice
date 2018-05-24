package main

import (
	"net"
	"fmt"
)

func main() {

	listen, err := net.Listen("tcp","0.0.0.0:50002")
	if err != nil {
		fmt.Println("listen tcp error:", err)
		return
	}

	for {
		connect, err := listen.Accept()
		if err != nil{
			fmt.Println("accept error:", err)
			continue
		}

		go handleConnect(connect)
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
	}
}

