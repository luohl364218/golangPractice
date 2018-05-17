//01 go语言关键字 标识符

//程序所属包
package main
//import printer "fmt"
import (//. "fmt"
	_ "golangPractice/practice00/learn02"
	"fmt"
)

//常量定义  首字母用大写
const NAME  = "heylink\n"
//全局变量
var mainName = "main name\n"
//全局变量
var a string = "nihao\n"
//一般声明
type myInt int
//结构的声明
type Learn struct {

}
//声明接口
type iLearn interface {

}
//函数定义
func learn()  {
	//printer.Print("test\n")
}
//常量显式定义
const golang string = "golang"
//常量隐式定义
const goland  = "goland"
//组合定义
const (
	cat string = "猫"
	dog = "狗"
)
//单行定义多个常量
const apple, banana string = "苹果", "香蕉"
const b, c  = 1, "测试"
const cLen  = len(c)
//自定义函数 含返回类型
func getString() string {
	return "getString"
}
//const getStringLen  = len(getString())   //错误 函数返回值不是const

const value1  = iota
//const value2  = iota

const (
	value2 = iota
	_ //跳值使用法
	_
	//value2_5 = 1.5  //插队使用法 value3+1
	value3 = iota
	value4 = iota * 2
	//表达式隐式使用法
	value5
	value6
	//单行使用法
	value7, value8 = iota, iota + 3
	value9, value10
	value11 = iota
)

func main() {
	/*********learn01***********/
	//learn()
	//fmt.Print("hello learn01\n")
	//fmt.Print(mainName)
	//fmt.Print(NAME)

	/*********learn02***********/
	/*验证import的加载原理
	import 的基本原理：
  	1.如果一个main导入其他包，包将被顺序导入
  	2.如果导入的包中依赖其他包（包B），会首先导入B包，然后初始化B包中的常量与变量，最后如果B中有init，会自动执行ini（）；
  	3.所有包到入完成之后才会对main中常量和变量进行初始化，然后执行main中的init函数（如果存在），最后执行main函数；
  	4.如果一个包被导入多次则该包只会被导入一次；

	注意：我在这里遇到坑
	以下两个方法调用一直报错，就是没有办法导入learn02和show02两个包
	原因是我的GOPATH路径下没有设置src文件夹，然后还要把当前项目放在src文件夹下
	go默认到GOPATH的src路径下去找源码，所以环境还是很重要的
	*/
	//show02.Show02()
	//learn02.Learn02()

	/*********learn03***********/
	/*
	import别名，“.”,"_"
	别名操作的含义是：将导入的包命名为另一个容易记忆的别名
	点（.）操作的含义是：点（.）标识的包导入后，调用该包中函数时可以省略前缀包名；
	下划线（_）操作的含义是：导入该包，但不导入整个包，而是执行该包中的init函数，
	因此无法通过包名来调用包中其他的函数。使用下划线（_）操作往往是为了注册包里的引擎，让外部可以方便地使用；
	*/
	//printer.Print("this is a alias for fmt!")
	//Print("this is a . for fmt!")

	/*********learn04***********/
	/*常量和变量*/
	//orange := "橘子"

	//fmt.Print(cat)
	//fmt.Print(dog)
	//fmt.Print(apple)
	//fmt.Print(banana)
	//fmt.Print(orange)
	//fmt.Print("\n")
	//fmt.Print(reflect.TypeOf(b))
	//fmt.Print("\n")
	//fmt.Print(b)
	//fmt.Print(c)
	//fmt.Print(cLen)
	/*********learn05***********/
	/*go语言的iota关键字
	iota在const关键字出现时将被重置为0
	const中新增一行常量声明将使iota计数一次
	iota常见使用法：
	1，跳值使用法；
	2，插队使用法；
	3，表达式隐式使用法；
	4，单行使用法。
	*/
	/*fmt.Print("value1:")
	fmt.Print(value1)
	fmt.Print("\n")
	fmt.Print("value2：")
	fmt.Print(value2)
	fmt.Print("\n")
	fmt.Print("value3：")
	fmt.Print(value3)
	fmt.Print("\n")
	//表达式隐式使用法
	fmt.Print("value4：")
	fmt.Print(value4)
	fmt.Print("\n")
	fmt.Print("value5：")
	fmt.Print(value5)
	fmt.Print("\n")
	fmt.Print("value6：")
	fmt.Print(value6)
	fmt.Print("\n")
	//单行使用法
	fmt.Print("value7：")
	fmt.Print(value7)
	fmt.Print("\n")
	fmt.Print("value8：")
	fmt.Print(value8)
	fmt.Print("\n")
	fmt.Print("value9：")
	fmt.Print(value9)
	fmt.Print("\n")
	fmt.Print("value10：")
	fmt.Print(value10)
	fmt.Print("\n")
	fmt.Print("value11：")
	fmt.Print(value11)
	fmt.Print("\n")*/

	/*********learn06***********/
	/*运算符等*/
	//只有a++，没有++a
	value := 0
	value++
	//fmt.Print(value)
	//if-else
	if value > 1 {
		fmt.Print("value > 1\n")
	} else if value < 7{
		fmt.Print("value < 7\n")
	} else {
		fmt.Print("value >= 7\n")
	}
	//switch
	switch value {
	case 1:
		fmt.Println("value = 1")
	case 2:
		fmt.Println("value = 2")
	default:
		fmt.Println("value = 3")
	}
	//for
	/*for {
		fmt.Println("死循环")
	}*/

	for i:=1; i < 10; i++{
		fmt.Println("for...")
		//time.Sleep(1*time.Second)
	}
	array := [] string{"苹果", "香蕉", "雪梨"}
	for key, valueNum := range array {
		fmt.Print("key的值为：")
		fmt.Println(key)
		fmt.Print("value的值为：")
		fmt.Println(valueNum)
	}
	//只取得value
	for _, valueNum := range array {
		fmt.Print("value的值为：")
		fmt.Println(valueNum)
	}

	//goto
	goto One
	fmt.Println("中间代码块")
	One:
		fmt.Println("goto代码块")

}

