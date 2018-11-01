package main

import "fmt"

func main() {

	var arr = [6]int{9, 8, 7, 6, 5, 4}
	var s1 = arr[2:5]

	fmt.Println(s1)

	fmt.Println(sun(s1))

	//var slice1 []int = make([]int, 10)
	//用 make() 创建一个切片,10是长度
	s2 := make([]int, 10)
	for i := 0; i < len(s2); i++ {
		s2[i] = 5 + i
	}
	fmt.Println(s2)

	//将切片扩展 1 位可以这么做：
	/*s4 := s2[0:len(s2)+1]
	fmt.Println(s4)*/

	//追加
	s3 := append(s2, 0, 1)
	fmt.Println(s3)

	/**
		func copy(dst, src []T) int copy 方法将类型为 T 的切片从源地址 src 拷贝到目标地址
	dst，覆盖 dst 的相关元素，并且返回拷贝的元素个数。源地址和目标地址可能会有重叠。拷
	贝个数是 src 和 dst 的长度最小值。如果 src 是字符串那么元素类型就是 byte。如果你还想继
	续使用 src，在拷贝结束后执行 src = dst 。
	*/

	s := "zhaojun"
	sub := s[0:len(s)]
	//获取字符串的第一个字符
	fmt.Println(sub[0])

	//修改字符串的第一个字符...
	c := []byte(s)
	c[0] = 'c'
	fmt.Println(string(c))

}

//将切片传递给函数
func sun(a []int) (res int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

/**
new(T) 为每个新的类型T分配一片内存，初始化为 0 并且返回类型为*T的内存地址：这
种方法 返回一个指向类型为 T，值为 0 的地址的指针，它适用于值类型如数组和结构体
；它相当于 &T{} 。
make(T) 返回一个类型为 T 的初始值，它只适用于3种内建的引用类型：切片、map 和
channel

new 函数分配内存，make 函数初始化；
*/
