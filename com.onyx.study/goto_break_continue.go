package main

import "fmt"

func main() {

	//goto  值跳转到执行位置

	//break  退出循环

	//continue  结束当前循环,继续下一次循环

	goto one
	fmt.Println("hah")
one:
	fmt.Println("goto")

	//goto one  //这就是无限循环了

}
