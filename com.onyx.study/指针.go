package main

import "fmt"

func main()  {

	var i=5;
	var i2=5;
	fmt.Println("i的地址是:",&i)
	fmt.Println("i2的地址是:",&i2)


	var intP *int
	intP=&i
	fmt.Println("指针的地址是:",intP)
	fmt.Println("这个intP对应的值是:",*intP)


	//改变后,指针没变,但是值变化了
	*intP=6
	fmt.Println("改变后,指针的地址是:",intP)
	fmt.Println("改变后,这个intP对应的值是:",*intP)
	fmt.Println(i)

	//不能得到一个文字或常量的地址
	//指针的一个高级应用是你可以传递一个变量的引用（如函数的参数），这样不会传递变量的拷贝


}
