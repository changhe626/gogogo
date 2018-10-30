package main

import "fmt"



func add2(a int,b int) (int,int){
	return a+1,b+1
}


func say(lists...string){
	fmt.Println(lists)

	defer fmt.Println("这个defer是用于资源的释放的...,在函数被返回前执行")
	fmt.Println("最后一句")
}


func fibonacci(n int)(res int){
	if n<=1{
		res=1
	}else {
		res=fibonacci(n-1)+fibonacci(n-2)
	}
	return res
}


func MyAdd(a,b int) {
	fmt.Println(a,b)
}

func callback(a int,f func(int,int)){
	f(a,a)
}


func main(){

	i2, i3 := add2(2, 4);

	fmt.Println(i2,i3)

	say("1","2","4","5")
	say()


	result:=0
	for i:=2;i<6 ;i++  {
		result=fibonacci(i)
		fmt.Println("i是",i,"的时候,结果是:",result)
	}


	//传递函数作为参数...
	callback(3,MyAdd)

}




