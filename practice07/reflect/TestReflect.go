package main

import (
	"reflect"
	"fmt"
)

type Student struct {
	Name string `heylink:"hahaha"`
}

func (s *Student) Print() {
	fmt.Println("this is a student:", s.Name)
}

func main() {
	var a int = 200
	testReflect(a)

	var stu = Student{
		Name:"heylink",
	}
	testReflect(stu)

	testStruct(&stu)

	b := 200
	//如果要更改值，传地址进去
	testInt(&b)
	fmt.Println(b)
}

func testReflect(b interface{})  {
	//获取类型
	t := reflect.TypeOf(b)
	fmt.Println(t)    //main.Student
	//获取值
	v := reflect.ValueOf(b)
	fmt.Println(v)    //{heylink}
	//获取类别
	k := v.Kind()
	fmt.Println(k)    //struct
	//把得到的value转回interface{}
	iv := v.Interface()
	stu, ok := iv.(Student)
	if ok {
		fmt.Printf("%v  %T\n", stu, stu)
	}
}

func testInt(b interface{}) {
	val := reflect.ValueOf(b)
	fmt.Printf("before set: %d\n", val.Elem())
	//Elem()方法相当于*
	val.Elem().SetInt(100)

	c:= val.Elem().Int()
	fmt.Printf("after set: %d\n", c)
}

//通过反射操作结构体
func testStruct(b interface{}) {
	val := reflect.ValueOf(b)
	kd := val.Kind()
	if kd != reflect.Ptr || val.Elem().Kind() == reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	//获取字段数量
	num := val.Elem().NumField()
	//修改字段内容
	val.Elem().Field(0).SetString("hexing")
	fmt.Printf("struct has %d fields\n", num)
	//获取方法数量
	num = val.NumMethod()
	fmt.Printf("struct has %d methods\n", num)
	//调用方法
	if num > 0 {
		val.Method(0).Call(nil)
	}
	//通过反射获取tag 要通过type拿
	st := reflect.TypeOf(b)
	field := st.Elem().Field(0)
	fmt.Println("反射获取tag：",field.Tag.Get("heylink"))
}