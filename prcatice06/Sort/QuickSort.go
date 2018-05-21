package main

import "fmt"

func main() {
	TestQuickSort()
}

func TestQuickSort()  {
	a := [...]int{9,8,7,4,5,2,1,3}
	quickSort(a[:], 0, len(a) - 1)
	fmt.Println(a)
}

func quickSort(a []int, left, right int)  {
	//递归退出条件
	if left >= right {
		return
	}

	//val所在的位置
	k := left
	val := a[left]
	//进行排序，比val小的都在左边，比val大的都在右边
	for i := left + 1; i <= right; i++{
		if a[i] < val {
			a[k] = a[i]
			a[i] = a[k + 1]
			k++
		}
	}
	a[k] = val

	//左边递归
	quickSort(a, left, k - 1)
	//右边递归
	quickSort(a, k + 1, right)
}
