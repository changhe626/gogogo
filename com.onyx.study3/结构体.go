package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	id   int
	age  int
	name string
}

func changeName(a *Animal) {
	a.name = strings.ToUpper(a.name)

}

func main() {

	cat := new(Animal)
	cat.name = "tom"
	cat.age = 2
	fmt.Println(cat.name)

	cat2 := Animal{1, 2, "jerry"}
	fmt.Println(cat2)

	var i1 Animal
	var i2 *Animal

	fmt.Println(i1.name)
	i2 = cat
	fmt.Println(i2.age)

	//一般都是传递指针,不会去传递一个对象,太大了.
	changeName(&cat2)
	fmt.Println("改变大小写之后:", cat2.name)

}
