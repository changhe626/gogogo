package main

import "fmt"

func add2(a int, b int) (int, int) {
	return a + 1, b + 1
}

func say(lists ...string) {
	fmt.Println(lists)

	defer fmt.Println("这个defer是用于资源的释放的...,在函数被返回前执行")
	fmt.Println("最后一句")
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return res
}

func MyAdd(a, b int) {
	fmt.Println(a, b)
}

func callback(a int, f func(int, int)) {
	f(a, a)
}

func main() {

	i2, i3 := add2(2, 4)

	fmt.Println(i2, i3)

	say("1", "2", "4", "5")
	say()

	result := 0
	for i := 2; i < 6; i++ {
		result = fibonacci(i)
		fmt.Println("i是", i, "的时候,结果是:", result)
	}

	//传递函数作为参数...
	callback(3, MyAdd)

}

/**
内置函数

close 用于管道通信

len、 cap
len 用于返回某个类型的长度或数量（字符串、 数组、 切片、 map 和管
道） ；cap 是容量的意思， 用于返回某个类型的最大容量（只能用于切片和 map）

new、make
new 和 make 均是用于分配内存：new 用于值类型和用户定义的类型， 如
自定义结构， make 用户内置引用类型（切片、 map 和管道） 。 它们的用
法就像是函数， 但是将类型作为参数：new(type)、 make(type)。 new(T)
分配类型 T 的零值并返回其地址， 也就是指向类型 T 的指针（详见第 10.1
节） 。 它也可以被用于基本类型： v := new(int) 。 make(T) 返回类型 T
的初始化之后的值， 因此它比 new 进行更多的工作 new() 是一个函数， 不要忘记它的括号

copy、append 用于复制和连接切片

panic、recover
两者均用于错误处理机制

print、println 底层打印函数（ ， 在部署环境中建议使用 fmt 包

complex、real imag 用于创建和操作复数（
*/
