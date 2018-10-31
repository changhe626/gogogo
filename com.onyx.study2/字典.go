package main

import "fmt"

func main() {

	var map1 = map[string]string{"a": "a", "b": "b"}

	mm := map[int]string{1: "a", 2: "b", 3: "c"}

	//运用索引表达式取出字典中的值
	//对于字典值来说，如果其中不存在索引表达式欲取出的键值对，
	// 那么就以它的值类型的空值（或称默认值）作为该索引表达式的求值结果。由于字符串类型的空值为""
	fmt.Println(map1["3"])
	fmt.Println(mm[2])

	delete(map1, "a")
	fmt.Println(map1)

	//与切片类型相同，字典类型属于引用类型。它的零值即为nil。

}
