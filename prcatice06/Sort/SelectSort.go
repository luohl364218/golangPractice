package main

import "fmt"

func main() {
	TestSelectSort()
}

func TestSelectSort()  {
	a := [...]int{9,8,7,4,5,2,1,3}
	selectSort(a[:])
	fmt.Println(a)
}

func selectSort(a []int)  {
	for i := 0; i < len(a); i++ {
		//每次选出最小的数
		for j := i+ 1;j < len(a); j++ {
			if a[j] < a[i] {
				a[j], a[i] = a[i], a[j]
			}
		}
	}
}
