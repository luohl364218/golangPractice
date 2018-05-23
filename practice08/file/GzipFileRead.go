package main

import (
	"os"
	"fmt"
	"compress/gzip"
	"bufio"
)
//读取压缩文件
func main() {
	fName := "myFile.gz"
	fi, err := os.Open(fName)
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}
	defer fi.Close()

	fz, err := gzip.NewReader(fi)
	if err != nil {
		fmt.Println("new reader error:", err)
		return
	}

	reader := bufio.NewReader(fz)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}

