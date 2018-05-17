package main

import (
	"golangPractice/practice01"
	"fmt"
	"golangPractice/practice02"
)

func main()  {
	fmt.Println("************practice01****************")
	practice01.GetOSInfo()
	fmt.Println("************practice02****************")
	//practice02.TestValueAndRefType()
	//practice02.TestGoRoute()
	practice02.TestPipe()
}
