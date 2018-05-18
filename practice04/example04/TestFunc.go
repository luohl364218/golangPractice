package main

import "fmt"

//测试函数的使用 函数在go语言中是一等公民
type op_func func(int, int) int

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

//自定义的函数也是一种类型
func operator(op op_func, a, b int) int {
	return op(a, b)
}

//也可以直接在参数中定义
func operator1(op func(int, int) int, a, b int) int {
	return op(a, b)
}

func main() {

	c := add
	sum := c(200, 300)
	fmt.Println(sum)

	result := operator(c, 200, 300)
	fmt.Println(result)

	//defer一般用来释放锁资源、关闭文件、释放连接等操作
	//演示defer关键字  当函数返回时会执行defer语句 所以一般用来做资源清理
	defer fmt.Println("defer result:", result)
	//多个defer语句，先进后出原则
	defer fmt.Println("defer sum:", sum)

	result = operator(sub, 200, 300)
	fmt.Println(result)

	result = multiply(30, 20)
	fmt.Println(result)

	result = addElements(10,2,54,56,100)
	fmt.Println(result)

	//匿名函数
	result = anonymous(100, 300)
	fmt.Println(result)
}

//给返回值定义名称  可以直接return
func multiply(a, b int) (c int) {
	c = a * b
	return
}

//可变参数  支持一个或多个参数相加
func addElements(a int, arg...int) int {
	var sum int = a
	for i := 0; i < len(arg); i++ {
		sum += arg[i]
	}
	return sum
}

//匿名函数
func anonymous(a, b int) int {
	//定义函数类型
	func1 := func(a1 int, b1 int) int{
		return a1 + b1
	}
	//调用函数执行
	result1 := func1(a, b)

	//定义匿名函数并调用
	result2 := func(a1 int, b1 int) int{
		return a1 + b1
	}(a, b)

	return result1 + result2
}