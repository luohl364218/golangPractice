package practice02

import "fmt"

//值类型 var
//引用类型ref 指针，slice，map，chan

func TestValueAndRefType()  {
	var a = 100
	var b chan int = make(chan int, 1)
	fmt.Println("a=", a)
	fmt.Println("b=",b)

	modify(a)
	fmt.Println("a=", a)

	modifyPoint(&a)
	fmt.Println("a=", a)
}

func modify(a int) {
	a = 10
	return
}

func modifyPoint(a *int) {
	*a = 10
}