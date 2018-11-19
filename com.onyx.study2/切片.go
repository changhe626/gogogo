package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

/**
切片（Slice）与数组一样，也是可以容纳若干类型相同的元素的容器。
与数组不同的是，无法通过切片类型来确定其值的长度。每个切片值都会将数组作为其底层数据结构。
切片是一个 长度可变的数组
*/
func main() {

	//注意，被“切下”的部分不包含元素上界索引指向的元素。
	// 另外，切片表达式的求值结果会是切片类型的，且其元素类型与被“切片”的值的元素类型一致。
	var num = [5]int{1, 2, 3, 4, 5}
	var slice1 = num[1:3] //2,3
	fmt.Println(reflect.TypeOf(slice1))

	//一个切片值的容量即为它的第一个元素值在其底层数组中的索引值与该数组长度的差值的绝对值。
	//为了获取数组、切片或通道类型的值的容量，我们可以使用内建函数cap
	fmt.Println(slice1, cap(slice1))

	//最后，要注意，切片类型属于引用类型。它的零值即为nil，即空值。
	// 如果我们只声明一个切片类型的变量而不为它赋值，那么该变量的值将会是nil
	var sli []int //不需要长度
	fmt.Println(sli, cap(sli))

	//这第三个正整数被称为容量上界索引。它的意义在于可以把作为结果的切片值的容量设置得更小。
	// 换句话说，它可以限制我们通过这个切片值对其底层数组中的更多元素的访问。
	var sli2 = num[1:4:4]
	fmt.Println(sli2)

	//append会对切片值进行扩展并返回一个新的切片值
	var sli3 = append(slice1, 5, 6, 9)
	fmt.Println(sli3)

	/**
		该操作的实施方法是调用copy函数。该函数接受两个类型相同的切片值作为参数，并会把第二个参数值中的元素复制到第一个参数值中的相应位置（索引值相同）上。这里有两点需要注意：

	  1. 这种复制遵循最小复制原则，即：被复制的元素的个数总是等于长度较短的那个参数值的长度。
	  2. 与append函数不同，copy函数会直接对其第一个参数值进行修改。
	*/

	fmt.Println("~~~~~~~~~")
	s := []int{1, 4, 6}

	fmt.Println(sum(s))

	fmt.Println("~~~~~~~~~")
	//slice1 := make([]type, len)
	s1 := make([]int, 10)
	fmt.Println(s1)

	//如果你想创建一个 slice1，它不占用整个数组，而只是占用以 len 为个数个项，那么只要
	//make 的使用方式是： func make([]T, len, cap)  ，其中 cap 是可选参数
	//make([]int, 50, 100)
	//new([100]int)[0:50]

	//进行字符串的拼接
	var b bytes.Buffer
	for i := 0; i < 5; i++ {
		b.WriteString(strconv.Itoa(i))
	}
	fmt.Println(b)

}

//将切片传递给函数
func sum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum

}

/**
new(T) 为每个新的类型T分配一片内存，初始化为 0 并且返回类型为*T的内存地址：这
种方法 返回一个指向类型为 T，值为 0 的地址的指针，它适用于值类型如数组和结构体；它相当于  &T{}  。
make(T) 返回一个类型为 T 的初始值，它只适用于3种内建的引用类型：切片、map 和
channel。
new 函数  分配内存，make 函数  初始化
*/
