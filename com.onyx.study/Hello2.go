package main

import sys"fmt"

func main(){
	sys.Println("导入的是包的别名")


	sys.Print(add(3,5))
}


//函数
func add(a int,b int) int {
	return a+b
}
//注释
/**
注释
 */

const  i  =3.24

const str string  ="zhaojun"

const (
	Unkonw=0
	Kuow=1
	)

var a,b *int
/**
当一个变量被声明之后，系统自动赋予它该类型的零值：int 为 0，float 为 0.0，bool 为 false，string 为空字符串，
指针为 nil。记住，所有的内存在 Go 中都是经过初始化的。

变量的命名规则遵循骆驼命名法，即首个单词小写，每个新单词的首字母大写，例如：numShips 和 startDate。
 */


 var a1=10
 var b1 bool=false







