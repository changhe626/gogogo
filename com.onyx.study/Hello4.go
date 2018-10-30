package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main(){

	/**
	产生随机数
	 */
	for i:=0;i<10 ;i++  {
		a:=rand.Int()
		fmt.Println(a)

	}

	/**
	产生一个范围内的随机数
	 */
	for i:=5;i>0 ;i--  {
		a:=rand.Intn(8)
		fmt.Println(a)
	}


	fmt.Println(strings.Contains("abc","12"))
	i:=strings.Index("abc","8")
	fmt.Println(i)




}

