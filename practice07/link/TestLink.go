package main

import "fmt"

func main() {
	var intLink Link = Link{}
	for i := 0; i < 10; i++ {
		//intLink.InsertHead(i)
		intLink.InsertTail(fmt.Sprintf("node：%d",i))
	}
	intLink.Traverse()
}
