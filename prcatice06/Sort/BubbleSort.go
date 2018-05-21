package main

import "fmt"

func main() {
	TestBubbleSort()
}

func TestBubbleSort()  {
	a := [...]int{9,8,7,4,5,2,1,3}
	bubbleSort(a[:])
	fmt.Println(a)
}

func bubbleSort(a []int)  {
	for i := 0; i < len(a); i++ {
		//每次冒泡排序固定最右端的数
		for j := 1;j < len(a) - i; j++ {
			if a[j] < a[j - 1] {
				a[j], a[j - 1] = a[j - 1], a[j]
			}
		}
	}
}
