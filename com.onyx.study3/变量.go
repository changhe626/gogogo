package main

import (
	"fmt"
	"reflect"
)

//变量可见性
//不可见
var b int = 1

//可见...
var NAME = "String"

func main() {

	//_  就是垃圾桶
	var a, _, c int = 1, 2, 3

	fmt.Println(a, c)
	// fmt.Println(_)   报错,打印不出来

	//类型转换,会有精度的损失....
	var a1 int = 2
	var b1 float32 = 3.01
	c1 := float32(a1)
	fmt.Println(c1, reflect.TypeOf(c1))
	a2 := int(b1)
	fmt.Println(a2, reflect.TypeOf(a2))

	//变量可见性
	fmt.Println(b)
	fmt.Println(NAME)

}
