package main

import "fmt"

func main() {
	ch := make(chan int,10)
	exitCh := make(chan struct{}, 2)
	go send(ch, exitCh)
	go recv(ch, exitCh)

	//等待所有goroute退出【采用计数方法比较稳】
	for i := 0; i < 2; i++{
		<- exitCh
	}
}

//应用场景：只写  chan<- int
func send(ch chan<- int, exitCh chan struct{})  {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("generate ", i)
	}
	//记得关闭
	close(ch)
	exitCh <- struct{}{}
}

//应用场景：只读，防止误操作  <-chan
func recv(ch <-chan int, exitCh chan struct{}) {
	for {
		v, ok := <- ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	exitCh <- struct{}{}
}
