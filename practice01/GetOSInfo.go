package practice01

import (
	"os"
	"fmt"
)
//获取操作系统名称和path环境变量
func GetOSInfo()  {
	var goos string= os.Getenv("GOOS")
	fmt.Printf("the operating system is: %s \n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("the os path is: %s \n", path)
}


