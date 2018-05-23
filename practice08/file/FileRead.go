package main

import (
	"os"
	"fmt"
	"bufio"
)

func main()  {
	//全路径
	file, err := os.Open("c:/keyvalue.txt")
	if err != nil {
		fmt.Println("read file err:",err)
		return
	}
	//关闭文件
	defer file.Close()

	reader := bufio.NewReader(file)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read string err:", err)
		return
	}

	slice := []rune(str)

	for _,v := range slice{
		fmt.Printf("%c", v)
	}

}

