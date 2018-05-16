//01 go语言关键字 标识符

//程序所属包
package main
//导入依赖包
import "fmt"
//常量定义  首字母用大写
const NAME  = "heylink"
//全局变量
var mainName = "main name"
//全局变量
var a string = "nihao"
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
	fmt.Print("test")
}

func main() {
	learn()

	fmt.Print("learn1")
	fmt.Print(mainName)
	fmt.Print(NAME)
}

