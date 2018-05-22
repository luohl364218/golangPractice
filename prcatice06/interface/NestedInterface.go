package main

import "fmt"

//接口嵌套

type Reader interface {
	Read()
}

type Writer interface{
	Write()
}

type ReadWriter interface {
	Reader
	Writer
}

type File struct {
}

func (f *File) Read(){
	fmt.Println("read data")
}

func (f *File) Write() {
	fmt.Println("write data")
}

func TestRW(rw ReadWriter) {
	rw.Read()
	rw.Write()
}

func main()  {
	var f File
	//为什么传指针类型？因为是指针类型实现了接口方法，不是值类型
	TestRW(&f)
}