//常用结构体...一般按照下面的顺序

//所属于包
package main

//引入依赖包
import "fmt"

//定义常量,大写名字
const NAME = "zhaojun"

//在main外面定义的变量就是全局变量
//定义的全局变量
var s = "zhangke"

//一般类型声明
type innodInt int

//结构声明
type lean struct {
}

//接口声明,前面加一个小写的i
type ilearn interface {
}

//函数定义
func add2() {
	fmt.Print("add2")
}

//main函数是程序入口,可执行所必须的
func main() {
	fmt.Println("hello world")
	fmt.Print(NAME)
	fmt.Print(s)
}
