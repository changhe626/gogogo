package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		//睡眠1s
		fmt.Println("haha", i)
		time.Sleep(1 * time.Second)
	}

	//无限循环
	for {
		fmt.Println("...")
	}

	//遍历数组
	a := []string{"a", "b", "c"}
	for key, value := range a {
		fmt.Println(key, value)
	}

	//使用_ 来接收...
	for _, value := range a {
		fmt.Println(value)
	}

}
