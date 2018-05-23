package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	//os.O_APPEND 追加到文件末尾，否则都会覆盖
	file, err := os.OpenFile("output.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file failed:",err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	str := "hello heylink\n"
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}

