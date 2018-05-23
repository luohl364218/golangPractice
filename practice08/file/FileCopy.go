package main

import (
	"os"
	"fmt"
	"io"
)

func main()  {
	CopyFile()
}

func CopyFile() {
	copyFile("output.txt", "outputCpy.txt")
	fmt.Println("copy done!")
}

func copyFile(srcName, dstName string) (written int64, err error) {
	srcFile, err := os.Open(srcName)
	if err != nil {
		fmt.Println("open src file err:", err)
		return
	}
	defer srcFile.Close()

	dstFile, err := os.OpenFile(dstName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open dst file err:", err)
		return
	}
	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)
}

