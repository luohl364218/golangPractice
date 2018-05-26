package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	connect, err := net.Dial("tcp", "localhost:10000")
	if err != nil {
		fmt.Println("dial tcp server host err:", err)
		return
	}

	var userId int
	var password string
	fmt.Println("please input userId and password:")
	fmt.Scanf("%d %s", &userId, &password)

	//TCP必须保持连接
	err = login(connect, userId, password)
	if err != nil {
		fmt.Println("login failed, err:", err)
		return
	}

	defer connect.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		sendStr, _ := reader.ReadString('\n')
		sendStr = strings.Trim(sendStr, "\n\r")
		if sendStr == "Q" {
			return
		}

		_, err := connect.Write([]byte(sendStr))
		if err != nil {
			return
		}
	}
}
