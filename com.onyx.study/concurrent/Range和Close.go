package main

import "fmt"

func main() {
	c:=make(chan int,10)
	go fibonacci(cap(c),c)
	v,ok :=<-c
	fmt.Println("v:",v,"ok,",ok)
	for i:= range c  {
		fmt.Println(i)
	}

}

func fibonacci(n int,c chan int)  {
	x,y:=1,1
	for i:=0;i<n;i++{
		c<-x
		x,y=y,x+y
	}
	close(c)
}

/**
for i := range c 能够不断的读取channel里面的数据，直到该channel被显式的关闭。上面代码我们看到可以显
式的关闭channel，生产者通过关键字 close 函数关闭channel。关闭channel之后就无法再发送任何数据了，在消费
方可以通过语法 v, ok := <-ch 测试channel是否被关闭。如果ok返回false，那么说明channel已经没有任何数据
并且已经被关闭。

记住应该在生产者的地方关闭channel，而不是消费的地方去关闭它，这样容易引起panic
 */
