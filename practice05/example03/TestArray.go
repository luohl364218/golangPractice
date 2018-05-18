package main

import (
	"fmt"
)

func main() {

	fab(10)
	//testArray()
	testMultidimensionalArray()
}

func testArray()  {
	//数组的初始化方法
	var a1 [5]int = [5]int{1,2,3,4,5}
	var a2 = [5]int{1,2,3,4,5}
	var a3 = [...]int{11,22,33,44}
	var a4 = [...]int{1:100, 3:300}
	var a5 = [...]string{1:"hello", 3:"world"}

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a5)
}

func testMultidimensionalArray()  {

	var f [3][3]int = [...][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//遍历多维数组
	for _, val := range f{
		for _, val2 := range val{
			fmt.Print(val2, " ")
		}
		fmt.Println()
	}
	//fmt.Println(f)
}


//非递归输出斐波那契数列前n项
func fab(n int)  {

	var a [] uint64
	a = make([]uint64, n)

	a[0] = 1
	a[1] = 1

	for i := 2; i < n; i++ {
		a[i] = a[i - 1] + a[i - 2]
	}

	for _, val := range a{
		fmt.Println(val)
	}
}

