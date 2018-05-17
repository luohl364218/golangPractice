package main

import (
	"math/rand"
	"fmt"
	"time"
)

//随机生成一定范围内的数
func main()  {
	//修改种子
	rand.Seed(time.Now().Unix())
	//生成10个随机整数
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Println(a)
	}
	//生成0-100之间的10个随机数
	for i := 0; i < 10; i++ {
		a := rand.Intn(100)
		fmt.Println(a)
	}
	//生成10个随机浮点数
	for i := 0; i < 10; i++ {
		a := rand.Float32()
		fmt.Println(a)
	}
}