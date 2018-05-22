package main

import "fmt"

//断言
func TestAssert(a interface{}) {
	//a.(type)判断某个变量是否实现了指定接口
	b , ok:= a.(int)
	if !ok {
		fmt.Println("convert failed")
		return
	}
	b += 3
	fmt.Println(b)
}

func classifier(items ...interface{})  {
	for index, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("%d params is bool, %v\n", index, v)
		case int, int32, int64:
			fmt.Printf("%d params is int, %v\n", index, v)
		case float32, float64:
			fmt.Printf("%d params is float, %v\n", index, v)
		case string:
			fmt.Printf("%d params is string, %v\n", index, v)
		}
	}
}

func main()  {
	var b int
	TestAssert(b)
	//接口断言，判断类型
	classifier(26,2.8,true, "hhhhh")
	var c int
	//空接口 所有类型都实现了空接口
	var a interface{}
	a = c
	fmt.Println(a)
}

