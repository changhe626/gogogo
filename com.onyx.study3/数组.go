package main

import "fmt"

func main() {

	//var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}

	array := [3]int{7, 8, 9}
	x := Sum(&array)

	fmt.Println(x)

}

//但是这种方式用的少
func Sum(a *[3]int) (res int) {
	sum := 0
	for _, v := range a {
		fmt.Println(v)
		sum += v
	}
	return sum

}
