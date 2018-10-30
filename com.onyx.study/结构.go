package main

import "fmt"

func main(){

	bool1:=false
	if bool1 {
		fmt.Println("true")
	}else {
		println("false")
	}


	if val:=10;val<100{
		println("<100")
	}else {
		println(">100")
	}


	num:=100
	switch num {

	case 100:
		print(100)
	case 99:
		print("99")
	default:
		print("default")

	}

	for i:=5;i>0 ;i--  {
		fmt.Println("hello world,",i)
	}

	for i:=5;i>0 ;i--  {
		for j:=0;j<5 ;j++  {
			println("循环中变量的值是:",i,j)
		}
	}


	str:="hello world"

	for pos,char := range str{
		fmt.Printf("内容是:%d -- %c",pos,char)
	}


}
