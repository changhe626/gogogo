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

	//map1 := make(map[keytype]valuetype)  。
	map1 := make(map[string]int, 10)
	fmt.Println(map1)

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

	//map 类型的切片
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)

}
