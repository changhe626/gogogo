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

	fmt.Println("~~~~~")

	var arr1 [5]int
	//遍历1
	for i := 1; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}
	//遍历2
	for index, value := range arr1 {
		fmt.Println(index, value)
	}

	fmt.Println("~~~~~")

	//遍历3
	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}

	b := a
	b[0] = "zhaojun"
	//赋值后,a,b的值不再相同
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("~~~~~~~~~")
	//只做部分的赋值
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	fmt.Println(arrKeyValue)

	sum(&num)

}

func sum(a *[4]int) {
	sum := 0
	for _, value := range a {
		sum += value
	}
	fmt.Println(sum)

}
