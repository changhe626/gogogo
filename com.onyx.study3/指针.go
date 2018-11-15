package main

import "fmt"

func change(i *int) int {
	return *i + 1

}

func modift(sel []int) {
	sel[0] = 11

}

func main() {

	//指针变量的类型为 *T
	i := 5
	var a *int = &i
	fmt.Println(i, a, *a)
	//指针的解引用可以获取指针所指向的变量的值。将 a 解引用的语法是 *a

	//指针的零值是 nil。
	var b *string
	fmt.Println(b)

	i2 := change(a)
	fmt.Println(i2, i)

	//传递的时候不用数组,而是使用切片
	arr := [3]int{1, 2, 4}
	modift(arr[:])
	fmt.Println(arr)

}
