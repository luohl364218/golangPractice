package main

import (
	"fmt"
)

func main()  {
	testMakeInit()
	testDirectInit()
	testMap()
	//测试查找
	testQuery()
	//测试遍历
	testTraverse()
	//测试删除
	testDelete()
	//测试切片
	testSlice()
}

func testMakeInit()  {
	a := make(map[string]string, 10)
	a["1"] = "1"
	a["2"] = "2"
	a["3"] = "3"
	a["3"] = "3"
	fmt.Println(a)

}

func testDirectInit()  {

	var b map[string]string = map[string]string{
		"1":"1",
		"2":"2",
		"3":"3",
		"4":"4",
	}

	fmt.Println(b)
}

//value也是map的情况
func testMap()  {
	fmt.Println("-------------testMap-------------")
	a := make(map[string]map[string]string, 8)

	a["1"] = make(map[string]string,5)

	a["1"]["1"] = "1"
	a["1"]["2"] = "2"
	a["1"]["3"] = "3"
	a["1"]["4"] = "4"

	fmt.Println(a["1"])
	fmt.Println(a)
}

func testQuery()  {

	a := make(map[string]map[string]string, 100)
	modify(a)

	fmt.Println(a)
}

func modify(a map[string]map[string]string)  {

	_, ok := a["heylinlook"]
	if !ok {
		a["heylinlook"] = make(map[string]string)
	}
	a["heylinlook"]["password"] = "123"
	a["heylinlook"]["nickname"] = "baobao"

	return
}

func testTraverse()  {
	fmt.Println("-------------testTraverse-------------")
	a := make(map[string]map[string]string, 8)

	a["1"] = make(map[string]string,5)
	a["1"]["1"] = "1"
	a["1"]["2"] = "2"
	a["1"]["3"] = "3"
	a["1"]["4"] = "4"

	a["2"] = make(map[string]string,5)
	a["2"]["1"] = "1"
	a["2"]["2"] = "2"
	a["2"]["3"] = "3"
	a["2"]["4"] = "4"

	for k,v := range a{
		fmt.Println(k)
		for k1, v1 := range v{
			fmt.Println("\t", k1, " ", v1)
		}
	}
}

func testDelete()  {
	fmt.Println("-------------testDelete-------------")
	a := make(map[string]map[string]string, 8)

	a["1"] = make(map[string]string,5)
	a["1"]["1"] = "1"
	a["1"]["2"] = "2"
	a["1"]["3"] = "3"
	a["1"]["4"] = "4"

	a["2"] = make(map[string]string,5)
	a["2"]["1"] = "1"
	a["2"]["2"] = "2"
	a["2"]["3"] = "3"
	a["2"]["4"] = "4"

	fmt.Println("before felete")

	for k,v := range a{
		fmt.Println(k)
		for k1, v1 := range v{
			fmt.Println("\t", k1, " ", v1)
		}
	}

	//删除整个map内容只能for循环，或者重新make
	delete(a, "2")

	fmt.Println("after felete")

	for k,v := range a{
		fmt.Println(k)
		for k1, v1 := range v{
			fmt.Println("\t", k1, " ", v1)
		}
	}
}

func testSlice()  {

	var a []map[int]int
	a = make([]map[int]int, 5)

	if a[0] == nil {
		a[0] = make(map[int]int)
	}
	a[0][1] = 1

	fmt.Println(a)
}