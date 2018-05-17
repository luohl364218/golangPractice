package main

import "fmt"
//字符与字符串
func main()  {
	var str = "hello world\n"
	var str1 = `
	床前明月光，
	疑是地上霜。
	`
	fmt.Println(str)
	fmt.Println(str1)

	var b byte = 'c'
	fmt.Println(b)
	fmt.Printf("%c", b)
}