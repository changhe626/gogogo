package main

import (
	"fmt"
)

func main() {

	s := "内事不明问张昭,外事不明还问周瑜"
	for i, v := range s {
		//Unicode 字符会占用 2 个字节，有些甚至需要 3 个或者 4 个字节来进行表示
		fmt.Printf("%d:%c ", i, v)
	}

	//截取字符串
	i := s[22:]
	fmt.Println(i)

	// 修改字符串中的某个字符
	s2 := "hello world"
	c := []byte(s2)
	c[0] = 'i'
	fmt.Println(s2, string(c))

}
