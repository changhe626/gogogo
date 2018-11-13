package main

import (
	"fmt"
)

//定义接口
type SalaryCalculator interface {
	CalculateSalary() int
}

//结构体
type Parmanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

//实现接口
func (p Parmanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

//实现接口
func (c Contract) CalculateSalary() int {
	return c.basicpay
}

//计算的函数
func calAllMoney(person ...SalaryCalculator) int {
	sum := 0
	for _, v := range person {
		sum = sum + v.CalculateSalary()
	}
	return sum
}

//没有包含方法的接口称为空接口。空接口表示为 interface{}。由于空接口没有方法，因此所有类型都实现了空接口。
func nullInterface(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {
	pemp1 := Parmanent{1, 5000, 20}
	c1 := Contract{1, 3}
	c2 := Contract{2, 4}
	money := calAllMoney(pemp1, c1, c2)
	fmt.Println("all money is ", money)

	nullInterface(pemp1)
	nullInterface("1")
	nullInterface(1)

}
