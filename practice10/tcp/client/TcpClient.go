package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {

	connect, err := net.Dial("tcp", "localhost:50002")
	if err != nil {
		fmt.Println("dial tcp server host err:", err)
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



