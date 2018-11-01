package main

import "fmt"

/**
iota在const关键字出现时将重置为 0
const中每新增一行常量声明将使iota计数一次;

跳值使用法
插队使用法
表达式隐式使用法
单行使用法

同一行iota的值是不会加的...
a,b=iota,iota+1
*/

const a = iota
const b = iota

const (
	c = iota * 2
	d
	e
)

func main() {

	fmt.Println(a) //0
	fmt.Println(b) //0

	fmt.Println(c) //0
	fmt.Println(d) //1
	fmt.Println(e) //正常2

	//要求e=3 ,需要在e前面加入一个_
	//这个就是调值使用法
	//需要再加一,中间再加一个 _

	//把d修改为3.14赋值
	//但是e 还是4

	//c=iota*2  的时候下面不变,还是一样的
	//但是下面就是
	//c=iota*2
	//	d
	//	e
	//会继承前面的非空表达式(使用往上的最近的一个非空的表达式),不断的累乘

}
