package main

import "fmt"

func main() {
	var a int = 10
	//取地址
	fmt.Println(&a)

	//指针定义和赋值
	var p *int
	p = &a
	fmt.Println(*p)

	//修改指针对应的值
	*p = 100
	fmt.Println(a)

	//修改指针指向的值
	var b int = 999
	p = &b
	fmt.Println(*p)
}
