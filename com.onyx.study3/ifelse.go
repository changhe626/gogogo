package main

import (
	"fmt"
)

func main() {

	a := 1
	if a > 1 {
		fmt.Println(">1")
	} else {
		fmt.Println("<=1")
	}

	if a < 1 {
		fmt.Println(">1")
	} else if a < 10 {
		fmt.Println("1-10")
	} else {
		fmt.Println(">10")
	}

	switch a {
	case 2:
		fmt.Println("case 2")
	case 1:
		fmt.Println("case 1")
	default:
		fmt.Println("default")
	}

	//类型判断
	var b interface{}
	b = 32
	switch b.(type) {
	case int:
		fmt.Println("int")
	case float32:
		fmt.Println("float32")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("other")
	}

}
