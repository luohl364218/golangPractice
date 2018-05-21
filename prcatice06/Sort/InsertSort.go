package main

import "fmt"

func main() {
	TestInsertSort()
}

func TestInsertSort()  {
	a := [...]int{9,8,7,4,5,2,1,3}
	insertSort(a[:])
	fmt.Println(a)
}

func insertSort(a []int)  {
	for i := 1; i < len(a); i++ {
		//每次比较当前数与之前一位
		for j := i;j > 0; j-- {
			if a[j] < a[j - 1] {
				a[j], a[j - 1] = a[j - 1], a[j]
			}
		}
	}
}
