package main

import (
	"fmt"
)

func main() {

	//var mapAssigned map[string]int   这样声明的话,下面放值就会报错的
	mapAssigned := map[string]int{}
	/*mapLit := map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)*/
	mapAssigned["a"] = 1
	mapAssigned["b"] = 2

	fmt.Println(mapAssigned)
	fmt.Println(mapAssigned["a"])

	//key是否存在啊
	_, ok := mapAssigned["a"]
	fmt.Println(ok)

	if _, ok := mapAssigned["a"]; ok {
		// ...
	}

	//从 map1 中删除 key1：
	//直接 delete(map1, key1) 就可以。

	//遍历
	for key, value := range mapAssigned {
		fmt.Println(key, value)
	}

}
