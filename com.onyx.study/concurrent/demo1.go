package main

import (
	"fmt"
	"runtime"
)

func say(s string){
	for i:=0;i<5;i++{
		//runtime.Gosched()表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine
		runtime.Gosched()
		fmt.Println(s)
	}
}

func sum(a []int,c chan int){
	sum:=0
	for _,v :=range a  {
		sum+=v
	}
	// send sum to c
	c<-sum
	/**
channel通过操作符 <- 来接收和发送数据
ch <- v //  发送 v 到 channel ch.
v := <-ch //  从 ch 中接收数据，并赋值给 v
	 */
}


func main() {

	go say("1")
	go say("2")
	say("3")


	a:=[]int{1,2,4,5,6,6}
	//c:=make(chan int)  尺寸过小会报错
	c:=make(chan int,9)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	go sum(a,c)
	x:=<-c
	y:=<-c

	fmt.Println(x,y,<-c)
	/**
	默认情况下，channel接收和发送数据都是阻塞的，除非另一端已经准备好，这样就使得Goroutines同步变的更加的
简单，而不需要显式的lock。所谓阻塞，也就是如果读取（value := <-ch）它将会被阻塞，直到有数据接收。其
次，任何发送（ch<-5）将会被阻塞，直到数据被读出。无缓冲channel是在多个goroutine之间同步很棒的工具
	 */


}
