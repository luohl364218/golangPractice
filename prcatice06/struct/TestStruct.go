package main

import "fmt"

type Student struct {
	//公有变量
	Name string
	//私有变量
	age int
	score float32
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
}


