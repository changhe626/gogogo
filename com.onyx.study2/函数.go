package main

import "fmt"

func main() {
	//您也可以直接对匿名函数进行调用：
	c := func(x, y int) int {
		return x + y
	}(3, 4)

	fmt.Println(c)

	/**
		表示参数列表的第一对括号必须紧挨着关键字 func ， 因为匿名函数没有名称。 花括号 {}
	涵盖着函数体， 最后的一对括号表示对该匿名函数的调用。
	*/
	func() {
		sum := 0
		for i := 1; i < 100; i++ {
			sum += i
		}
	}()

	fmt.Println(f())

}

//变量 ret 的值为 2， 因此 ret++ ， 这是在执行 reutrn 1 语句后发生的
func f() (res int) {
	defer func() {
		res++
	}()

	return 1
}

/**
匿名函数同样被称之为闭包（函数式语言的术语） ：它们被允许调用定义在其它环境下的变
量。 闭包可使得某个函数捕捉到一些外部状态， 例如：函数被创建时的状态。 另一种表示方
式为：一个闭包继承了函数所声明时的作用域。 这种状态（作用域内的变量） 都被共享到闭
包的环境中， 因此这些变量可以在闭包中被操作， 直到被销毁，  闭包
经常被用作包装函数：它们会预先定义好 1 个或多个参数以用于包装， 详见下一节中的示
例。 另一个不错的应用就是使用闭包来完成更加简洁的错误检查
*/
