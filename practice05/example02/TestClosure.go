package main

import (
	"fmt"
	"strings"
)

func Adder() func(int) int {
	//有点像静态变量，会记住每一次调用后修改的最新值
	var x int
	return func(i int) int {
		x += i
		return x
	}
}

//闭包 一个函数和与其相关的引用环境组合而成的实体
func main() {

	//闭包
	f := Adder()
	fmt.Println(f(1))    //1
	fmt.Println(f(100))  //101
	fmt.Println(f(1000)) //1101

	f1 := makeSuffix(".jgp")
	fmt.Println(f1("boy"))
	fmt.Println(f1("girl"))

}

//经典例子
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}


