package main

import "fmt"

type Describer interface {
	Describe()
}

type Eat interface {
	EatSome() int
}

type Person struct {
	name string
	age  int
}

func (p Person) Describe() { // 使用值接受者实现
	fmt.Println(p.name, p.age)

}

//实现了多个接口
func (p Person) EatSome() int {
	fmt.Println(p.name, "is eating something")
	return 1
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describe() { // 使用指针接受者实现
	fmt.Println(a.country, a.state)
}

func main() {

	p1 := Person{"zhaojun", 1}
	p1.EatSome()

	a1 := Address{"aa", "bb"}
	p1.Describe()

	a1.Describe()

	var d1 Describer
	p2 := Person{"zk", 2}
	d1 = p2
	d1.Describe()

	var d2 Describer
	d2 = &p2
	d2.Describe()

	//接口的零值是 nil。对于值为 nil 的接口，其底层值（Underlying Value）和具体类型（Concrete Type）都为 nil。
	var d3 Describer
	fmt.Println(d3) //d3 就是个nil

}
