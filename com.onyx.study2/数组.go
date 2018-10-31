package main

import "fmt"

/**
一个数组（Array）就是一个可以容纳若干类型相同的元素的容器。
这个容器的大小（即数组的长度）是固定的，且是体现在数组的类型字面量之中的
*/
func main() {

	var num = [4]int{1, 1, 14, 4}
	fmt.Println(num[0], len(num))
	num[0] = 9
	fmt.Println(num[0], len(num))

}
