//01 go语言关键字 标识符

//程序所属包
package main
//import printer "fmt"
import (//. "fmt"
	_ "golangPractice/learn02"
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

}

