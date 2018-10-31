package main

import "fmt"

type Person struct {
	Name   string
	Gender string
	Age    uint8
}

func (person *Person) Grow() {
	person.Age++
}

func main() {

	p := struct {
		Name   string
		Gender string
		Age    uint8
	}{"Robert", "Male", 30}

	p1 := Person{Name: "Robert", Gender: "Male", Age: 33}

	p1.Grow()

	p2 := Person{"Robert", "Male", 34}

	fmt.Println(p)
	fmt.Println(p1)
	fmt.Println(p2)

}
