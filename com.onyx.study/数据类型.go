package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

//类型设置别名
type zhaojun int32

func main() {

	/**
	整形,浮点型,复数,字符串,布尔
	字符串同意为utf-8
	*/

	var i uint8 = 1
	//这个变量占用内存大小,1字节
	fmt.Println(unsafe.Sizeof(i))
	//4个字节
	var i2 int32 = 5
	fmt.Println(unsafe.Sizeof(i2))
	//8个字节,根据系统决定,64位是8字节,32位是4字节
	var i3 int = 5
	fmt.Println(unsafe.Sizeof(i3))

	//4
	var f1 float32 = 3.4
	fmt.Println(unsafe.Sizeof(f1))
	//8
	var f2 float64
	fmt.Println(unsafe.Sizeof(f2))

	//只能是true/false   1
	var b1 bool = false
	fmt.Println(unsafe.Sizeof(b1))

	//1
	var b2 byte = 1
	fmt.Println(unsafe.Sizeof(b2))

	//4
	var r1 rune
	fmt.Println(unsafe.Sizeof(r1))

	//派生类型...

	//类型零值,就是默认值, bool是false,数值是0,字符串是空字符串
	//类型设置别名
	var z zhaojun
	//获取  默认值  变量类型   占用空间大小     0 main.zhaojun 4
	fmt.Println(z, reflect.TypeOf(z), unsafe.Sizeof(z))

	//别名zhaojun的变量和int32不能进行计算...

}
