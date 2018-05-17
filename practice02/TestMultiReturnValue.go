package practice02

import "fmt"

func TestMultiReturnValue()  {
	sum, avg := calc(100, 200)
	fmt.Println("sum=", sum, " avg=", avg)
}

func calc(a int, b int) (int, int) {
	c := a + b
	d := (a + b) / 2
	return c, d
}
