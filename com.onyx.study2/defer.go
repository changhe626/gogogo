package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	//逆序执行
	for i := 0; i < 5; i++ {
		defer fmt.Println("i=", i)

	}

	/**
	关键字 defer 允许我们进行一些函数执行完成后的收尾工作， 例如：
	1. 关闭文件流：
	// open a file defer file.Close()
	2. 解锁一个加锁的资源
	mu.Lock() defer mu.Unlock()
	3. 打印最终报告
	printHeader() defer printFooter()
	4. 关闭数据库链接
	// open a database connection defer disconnectFromDB()
	合理使用 defer 语句能够使得代码更加简洁。
	*/

	//使用 defer 语句来记录函数的参数与返回值
	func1("gogogog")

}

func func1(s string) (n int, err error) {

	defer func() {
		log.Println(s, n, err)
	}()

	return 7, errors.New("cuowu")

}
