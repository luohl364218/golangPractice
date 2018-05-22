package main

import (
	"fmt"
	"time"
)

type Student struct {
	//公有变量
	Name string  "这个是tag,可以通过反射获取"
	//私有变量
	age int  "可以为每个属性添加tag"
	score float32
}

type Cart struct {
	//匿名字段
	Student
	int
	start time.Time
}

func main() {
	//初始化
	var stu Student
	stu.Name = "heylink"
	stu.age = 25
	stu.score = 99.9
	fmt.Println(stu)

	var stu1 *Student = &Student{
		age:25,
		Name:"hexing",
		score:99.9,
	}
	fmt.Println(*stu1)

	var stu2 *Student = new(Student)
	stu2.Name = "ff"
	stu2.score = 0.0
	stu2.age = 22
	fmt.Println(*stu2)

	var stu4 = Student{
		age:25,
		Name:"hhh",
		score:99.9,
	}
	fmt.Println(stu4)

	//观察内存空间分布
	fmt.Printf("name:%p\n", &stu.Name)
	fmt.Printf("age:%p\n",&stu.age)
	fmt.Printf("score:%p\n", &stu.score)

	//匿名字段
	var cart Cart
	cart.Name = "test"
	cart.age = 76
	cart.score = 1.1
	cart.int = 100

	fmt.Println(cart)
}

//结构体没有构造函数，需要的话自己定义
func NewStudent(name string, age int, score float32) *Student {
	return &Student{
		Name:name,
		age:age,
		score:score,
	}
}

