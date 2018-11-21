package main

/**
读写锁是针对于读写操作的互斥锁。
基本遵循两大原则：
1、可以随便读。多个goroutin同时读。
2、写的时候，啥都不能干。不能读，也不能写。


解释：
在32位的操作系统中，针对int64类型值的读操作和写操作不可能只由一个CPU指令完成。如果一个写的操作刚执行完了第一个指令，
时间片换给另一个读的协程，这就会读到一个错误的数据。


RWMutex提供四个方法：

func (*RWMutex) Lock //写锁定
func (*RWMutex) Unlock //写解锁
func (*RWMutex) RLock //读锁定
func (*RWMutex) RUnlock //读解锁
 */

import (
	"sync"
	"time"
)



var m *sync.RWMutex



func main() {

	m = new(sync.RWMutex)

	//可以多个同时读
	go read(1)
	go read(2)

	time.Sleep(2 * time.Second)

}



func read(i int) {
	println(i, "read start")
	m.RLock()
	println(i, "reading")
	time.Sleep(1 * time.Second)
	m.RUnlock()

	println(i, "read end")

}