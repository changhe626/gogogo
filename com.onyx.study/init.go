package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
)

var Pi float64

//先执行的init这个函数
func init() {

	Pi = 4 * math.Atan(1) // init() function computes Pi
}

func main() {
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)

	fmt.Println("hello world", 90)

	fmt.Println(Pi)

}
