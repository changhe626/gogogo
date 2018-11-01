package main

import "fmt"

//定义常量,常量只有 布尔,数字(整形,浮点,复数),字符串
const name = "zhaojun"

const name2 string = "zhaojun2"

//定义多个常量
const (
	name3     = "zhaojun3"
	name4 int = 4
)

//一次定义多个常量
const name5, name6 = "zhaojun5", "zhaojun6"
const name7, name8 string = "zhaojun5", "zhaojun6"

//使用表达式的方式,只能够使用内置的函数
const name9 = len(name)

func main() {
	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(name3)
	fmt.Println(name4)

	fmt.Println(name9)

}
