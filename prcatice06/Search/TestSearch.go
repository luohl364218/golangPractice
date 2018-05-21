package main

import (
	"fmt"
	"sort"
)

func main()  {
	testSearch()
}

func testSearch()  {

	var intArr = [...]int{39,23,2,3,54,26,42}
	//传入数组切片
	sort.Ints(intArr[:])

	index := sort.SearchInts(intArr[:], 23)
	fmt.Println(intArr)
	fmt.Println(index)
}