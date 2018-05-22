package main

import (
	"golangPractice/practice07/balance/source"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main()  {

	var insts []*source.Instance
	for i := 0; i < 10; i++ {
		ins := source.NewInstance(fmt.Sprintf("192.168.%d.%d",rand.Intn(255), rand.Intn(255)), rand.Intn(10000))
		insts = append(insts,ins )
	}

	//随机调度
	//balancer := &source.RandomBalance{}
	//轮询
	/*balancer := &source.RoundRobin{}

	for true {
		inst, _ := balancer.DoBalance(insts)
		fmt.Println(inst)
		time.Sleep(time.Second)
	}*/

	var balanceName = "random"
	if len(os.Args) > 1 {
		balanceName = os.Args[1]
	}

	for true {
		inst, _ := source.DoBalance(balanceName, insts)
		fmt.Println(inst)
		time.Sleep(time.Second)
	}

}
