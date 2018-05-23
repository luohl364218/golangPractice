package main

import (
	"fmt"
	"os"
	"time"
)

type ErrDiyType struct {
	path string
	createTime string
	message string
}

func (p *ErrDiyType)Error() string  {
	return fmt.Sprintf("ErrorInfo:\npath:%s\ncreateTime:%s\nmessage:%s\n",
		p.path, p.createTime, p.message)
}

func Open(name string) error {

	file, err := os.Open(name)
	if err != nil {
		return &ErrDiyType{
			path:"c:/test/error/",
			createTime:time.Now().Format("2006/01/02 15:04:05"),
			message:err.Error(),
		}
	}

	defer file.Close()
	return nil
}

func main() {
	err := Open("c:/abcd.txt")
	if err != nil {
		fmt.Println(err)
	}
}

