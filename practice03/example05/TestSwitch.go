package main

import (
	"fmt"
	"math/rand"
)

//switch语句中fallthrough穿透，没有break语句跳出，默认都是跳出
func main()  {

	a := 6
	//有判断条件
	switch a {
	case 0:
		fmt.Println("值为0")
		fallthrough
	case 1:
		fmt.Println("值为1")
		fallthrough
	case 2:
		fmt.Println("值为2")
	case 3, 4, 5, 6, 7, 8:
		fmt.Println("值在3, 4, 5, 6, 7, 8之间")
	default:
		fmt.Println("值未知")
	}

	//无条件判断
	switch {
	case a > 0 && a < 5:
		fmt.Println("a > 0 && a < 5")
	case a >= 5 && a < 10:
		fmt.Println("a >= 5 && a < 10")
	default:
		fmt.Println("default")
	}

	//语句块判断 后面要加分号
	switch a = 3 * a;{
	case a > 0 && a < 5:
		fmt.Println("a > 0 && a < 5")
	case a >= 5 && a < 10:
		fmt.Println("a >= 5 && a < 10")
	default:
		fmt.Println("a >= 10")
	}

	//判断有没有猜对随机数
	n := rand.Intn(100)
	for{
		var input int
		fmt.Scanf("%d", &input)
		flag := false

		switch {
		case input == n:
			flag = true
			fmt.Println("you are right!!!!!")
		case input > n:
			fmt.Println("bigger")
		case input < n:
			fmt.Println("less")
		}

		if flag {
			break
		}
	}

}
