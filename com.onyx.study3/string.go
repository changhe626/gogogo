package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	s := "昭君正在上自习"

	fmt.Println(strings.Contains(s, "昭君"))
	fmt.Println(strings.Compare("zk", "ZK"))
	fmt.Println(strings.Split("1,2,3,4", ","))

	//加入分割符
	fmt.Println(strings.Join([]string{"1", "2"}, "&"))

	//string 2 int
	fmt.Println(strconv.Atoi("4"))

}
