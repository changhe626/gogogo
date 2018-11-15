package main

import "fmt"

/**
所有信道都关联了一个类型。信道只能运输这种类型的数据，而运输其他类型的数据都是非法的。

chan T 表示 T 类型的信道。

信道的零值为 nil。信道的零值没有什么用，应该像对 map 和切片所做的那样，用 make 来定义信道。
*/
func main() {
	var a chan int
	if a == nil {
		a = make(chan int)
		fmt.Printf("Type of a is %T", a)
	}
	//data := <- a // 读取信道 a
	//a <- data // 写入信道 a

}
