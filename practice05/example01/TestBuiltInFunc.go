package main

import (
	"fmt"
	"errors"
)

//go语言内置函数
func main()  {

	//1.new 与直接定义的区别:new为值类型分配内存，返回指针
	var i int
	fmt.Println(i)

	j := new(int)
	fmt.Println(j)

	//2. make 为引用类型分配内存
	channel := make(chan int, 2)
	fmt.Println(channel)

	tellNewAndMake();

	//3. append 用来追加元素到数组， slice中  定义了具体长度的是数组，没有定义长度的是切片slice
	//var a [5]int
	var a []int
	a = append(a, 10, 20, 30)
	fmt.Println(a)
	//复制一份  ...除了可变参数外还可以展开数组
	a = append(a, a...)
	fmt.Println(a)

	//4. panic与recover：主要用作错误处理
	test()
	
}

//4. panic与recover：主要用作错误处理
func test() {

	defer func() {
		//recover捕获异常
		if err := recover();err != nil {
			fmt.Println(err)
		}
	}()

	//1.系统异常
	/*b := 0
	a := 100 / b
	fmt.Println(a)*/
	//2.自己抛出异常
	err := initConfig()
	if err != nil {
		//panic抛出异常
		panic(err)
	}
}

func initConfig() (err error){
	return errors.New("init config failed!")
}

func tellNewAndMake() {
	//new一个切片
	s1 := new([]int)
	fmt.Println(s1)  //&[]
	//初始化一个指针而已，未初始化实际的切片

	//make一个切片
	s2 := make([]int, 10)
	fmt.Println(s2)  //[0 0 0 0 0 0 0 0 0 0]
	//实际初始化了一个切片并赋初值

	*s1 = make([]int, 5)
	(*s1)[0] = 100
	s2[0] = 100
	fmt.Println(s1)  //&[100 0 0 0 0]
	return
}



