package main

import "fmt"

type Car struct {
	weight int
	name string
}

//父类方法也可以被继承
func (p *Car) Run()  {
	fmt.Println("running...")
}

//重写string接口 与java中的string相似
func (p Car) String() string {
	return fmt.Sprintf("weight=%d name=%s", p.weight, p.name)
}

type Vehicle struct {
	status bool
	running bool
}

type Bike struct {
	//通过匿名字段实现继承
	Car
	//多个匿名字段实现多重继承
	Vehicle
	wheel int
}

type People struct {
	//有名字的这种叫组合，所以说匿名字段是特殊的组合
	car Car
	house int
	name string
	age int
}

func main() {
	var a Bike
	//调用父类字段
	a.weight = 100
	a.name = "bike"
	a.wheel = 2
	a.status = true
	a.running = true

	fmt.Println(a)
	//调用父类方法
	a.Run()

	car := Car{
		weight:8,
		name:"car",
	}

	fmt.Println(car)
}

