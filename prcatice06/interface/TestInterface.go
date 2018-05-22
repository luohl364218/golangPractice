package main

import "fmt"

type Test interface {
	Print()
	//Write()
}

type Student struct {
	name string
	age int
	score int
}

//go中，任何类型 只要实现了某个接口中的所有方法，就实现了该接口
func (p *Student) Print()  {
	fmt.Println("name:", p.name)
	fmt.Println("age:", p.age)
	fmt.Println("score:", p.score)
}

func main() {

	var t Test
	var student Student = Student{
		name:"heylink",
		age:16,
		score:99,
	}
	t = &student
	t.Print()
}

