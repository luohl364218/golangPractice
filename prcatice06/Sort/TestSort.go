package main

import (
	"sort"
	"fmt"
)

func main (){
	//整数排序
	testIntSort()
}

func testIntSort() {
	var intArr = [...]int{39,23,2,3,54,26,42}
	//传入数组切片
	sort.Ints(intArr[:])
	fmt.Println(intArr)
}