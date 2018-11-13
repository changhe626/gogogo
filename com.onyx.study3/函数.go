package main

import "fmt"

func read(a, b int) (c, d int) {
	return 1 + 1, b + 2
}

func read2(a int, b int) (int, int) {
	return 1 + 1, b + 2
}

func read3(a ...int) {
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}

func main() {
	i, d2 := read(1, 4)
	fmt.Println(i, d2)

	i2, _ := read2(1, 3)
	fmt.Println(i2)

	read3(1, 2, 4)
	read3()
}
