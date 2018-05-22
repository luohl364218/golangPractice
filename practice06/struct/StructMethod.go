package main

import "fmt"

type Stud struct {
	name string
	age int
	score float32
}

//定义结构体的方法
func (this *Stud) init(name string, age int, score float32){
	//this不是指针的话，传入的是拷贝，无法真正改变原来的值
	this.name = name
	this.age = age
	this.score = score
	fmt.Println(this)
}

func (this Stud) get() Stud{
	return this
}

func (p Stud) getName() string {
	return p.name
}

func main() {

	var stu Stud
	//虽然init方法需要传入指针，但是go会自动帮你转
	stu.init("heylink", 34, 100)

	stu1 := stu.get()
	fmt.Println(stu1)

	name := stu.getName()
	fmt.Println(name)
}