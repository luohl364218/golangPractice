package practice02

import (
	"fmt"
	"time"
)

//管道是队列，先进先出,放不下的话会阻塞
func TestPipe() {
	fmt.Println("start goroute")
	//go testErrPipe()
	fmt.Println("start testPipeCommunication1")
	testPipeCommunication1()
	fmt.Println("start testPipeCommunication2")
	testPipeCommunication2()
	fmt.Println("end goroute")
	time.Sleep(time.Second)
}

func testPipe()  {
	//3是管道容量
	pipe := make(chan int, 3)
	pipe <- 1
	pipe <- 2
	pipe <- 3
	//pipe <- 4  //fatal error: all goroutines are asleep - deadlock!

	var t1 int
	//取出管道数据
	if len(pipe) > 0 {
		t1 = <- pipe
	}else {
		t1 = -1
	}

	pipe <- 4

	fmt.Println(t1)
	fmt.Println(len(pipe))
}

func testErrPipe()  {
	//3是管道容量
	pipe := make(chan int, 3)
	pipe <- 1
	pipe <- 2
	pipe <- 3
	pipe <- 4  //fatal error: all goroutines are asleep - deadlock!
	fmt.Println(len(pipe))
}
/**************管道通信***************/

/*——————全局变量——————————*/
//声明全局管道
var newPipe chan int
//管道通信方法1
func testPipeCommunication1()  {
	//创建管道
	newPipe = make(chan int, 1)
	go addNum(5, 3)
	//阻塞等待管道中的返回值
	fmt.Println("waiting for result.......")
	sum := <- newPipe
	fmt.Println("sum=", sum)
	//time.Sleep(time.Second)
}

func addNum(a int, b int) {
	var sum int = a + b
	time.Sleep(3 * time.Second)
	newPipe <- sum
}

/*——————管道作为变量传递——————————*/

func testPipeCommunication2()  {
	pipe := make(chan int, 1)
	go addNum2(5, 3, pipe)
	//阻塞等待管道中的返回值
	fmt.Println("waiting for result.......")
	sum := <- pipe
	fmt.Println("sum=", sum)
}

func addNum2(a int, b int, c chan int) {
	var sum = a + b
	time.Sleep(3 * time.Second)
	c <- sum
}
