package main

import "fmt"

//格式化输出
/*
https://go-zh.org/pkg/fmt/

一般：

%v	相应值的默认格式。在打印结构体时，“加号”标记（%+v）会添加字段名
%#v	相应值的Go语法表示
%T	相应值的类型的Go语法表示
%%	字面上的百分号，并非值的占位符
布尔：

%t	单词 true 或 false。
整数：

%b	二进制表示
%c	相应Unicode码点所表示的字符
%d	十进制表示
%o	八进制表示
%q	单引号围绕的字符字面值，由Go语法安全地转义
%x	十六进制表示，字母形式为小写 a-f
%X	十六进制表示，字母形式为大写 A-F
%U	Unicode格式：U+1234，等同于 "U+%04X"

*/
func main()  {

	var a int = 100
	var b bool
	var c byte = 'a'

	//%v保持原样
	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n",b)
	fmt.Printf("%v\n", c)

	//整数转字符串
	str := fmt.Sprintf("%d", a)
	fmt.Print(str)
	//含双引号
	fmt.Printf("%q\n", str)

	//字符串拼接
	str1 := "hello "
	str2 := "world"
	//str3 := str1 + str2 + "!"
	str3 := fmt.Sprintf("%s %s", str1, str2)
	len := len(str3)
	fmt.Print(str3)
	fmt.Println(" lenght:" , len)
	//切片
	substr := str3[2:5]
	fmt.Println(substr)

	substr = str3[6:]
	fmt.Println(substr)
	
	//字符串反转
	reverse := reverse(str3)
	fmt.Println(reverse)
	reverse = reverse1(reverse)
	fmt.Println(reverse)
}

//字符串反转
func reverse(str string) string {
	var result string
	strLen := len(str)

	for i := strLen - 1; i >= 0; i-- {
		result += fmt.Sprintf("%c", str[i])
	}
	return result
}

func reverse1(str string) string {
	var result []byte
	temp := []byte(str)
	strLen := len(str)
	for i := strLen - 1; i >= 0; i-- {
		result = append(result, temp[i])
	}
	return string(result)
}
