package main

import (
	"fmt"
	"reflect"
)

type File struct {
	fd   int
	name string
}

type TagType struct {
	/**
		标签（tag）：它是一个附属于字
	段的字符串，可以是文档或其他的重要标记。标签的内容不可以在一般的编程中使用，只有
	包  reflect  能获取它。
	*/
	id   bool   "这是id"
	name string "这是名字"
}

//匿名字段
//在一个结构体中对于每一种数据类型只能有一个匿名字段
type num struct {
	id int
	string
}

//下面是这个结构体类型对应的工厂方法，它返回一个指向结构体实例的指针
func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}

}

func main() {
	tt := TagType{true, "zhaojun"}
	for i := 0; i < 2; i++ {
		ref(tt, i)
	}

	n := new(num)
	n.id = 1
	n.string = "曹操"
	fmt.Println("匿名的对象是:", n)

}

func ref(tagType TagType, i int) {
	typeOf := reflect.TypeOf(tagType)
	field := typeOf.Field(i)
	fmt.Println(field.Tag)
}

/**
试图  make()  一个结构体变量，会引发一个编译错误，这还不是太糟糕，但是  new()  一个映
射并试图使用数据填充它，将会引发运行时错误！ 因为  new(Foo)  返回的是一个指向  nil
的指针，它尚未被分配内存。所以在使用  map  时要特别谨慎。
*/

/**
 当两个字段拥有相同的名字（可能是继承来的名字）时该怎么办呢？
1. 外层名字会覆盖内层名字（但是两者的内存空间都保留），这提供了一种重载字段或方
法的方式
2. 如果相同的名字在同一级别出现了两次，如果这个名字被程序使用了，将会引发一个错
误（不使用没关系）。没有办法来解决这种问题引起的二义性，必须由程序员自己修正。
*/
