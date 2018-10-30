package main

import "fmt"

//全局变量必须用var,局部变量可以省略

//特殊变量   _   不再需要的变量接收...

var a int
var b0 int = 4
var NAME string

func main() {

	var i int
	var f1 float32 = 43
	var i3 int = 4

	var (
		i4 string
		i5 bool
	)

	fmt.Println(i, f1, i3, i4, i5)
	fmt.Println(a)

	//局部覆盖全局变量...
	var a, b, c int = 1, 2, 3
	d, e := 1, 2

	fmt.Println(a)
	fmt.Println(b, c, d, e)

	//类型转换,必须是显示的转换
	//<变量名称> [:] =<目标类型> (<需要转换的变量>)
	f2 := float64(f1)

	fmt.Println(f2)

	//大写字母开头的变量是可导出的,也就是其他包可以读取的,是公用变量
	//小写字母开头是不可导出的,是私有变量

	//类型的自动推断,以下的简化方式
	var a, b, v int = 1, 2, 3
	var a, b, v = 1, 2, 3
	g, h := 1, 2

}
