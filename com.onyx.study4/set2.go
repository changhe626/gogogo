package main

/**
title:一个切片实现的不重复的Set
author:create by zk
time:2018年11月21日12:24:20
 */
import (
	"reflect"
	"fmt"
	"sync"
)

//这里使用互斥锁
var m *sync.Mutex= new(sync.Mutex)

/**
定义数据结构
 */
type set struct {
	arr []interface{}
}

/**
获取一个新的Set
 */
func  NewSet() *set{
	return &set{make([]interface{},0)}
}


/**
添加元素
 */
func (s *set)Add(v interface{}){
	m.Lock()
	//暂时先不考虑容量啥的问题
	flag:=s.Contains(v)
	if(!flag){
		s.arr=append(s.arr,v)
	}
	defer m.Unlock()
}

/**
添加一组元素
 */
func (s *set) AddAll (v...interface{}){
	for _,value:=range v{
		fmt.Println(value)
		s.Add(value)
	}
}

/**
是否包含
 */
func (s *set) Contains (v interface{}) bool {
	flag:=false
	for _,value :=range s.arr  {
		if(reflect.DeepEqual(value, v)){
			flag=true
		}
	}
	return  flag
}

/**
Set的尺寸
 */
func (s *set) Size() int{
	return len(s.arr)
}

/**
是否为空
 */
func (s *set) IsEmpty() bool {
	flag:=true
	if s.Size()>0{
		for _,value:= range s.arr{
			if(value!=nil){
				flag=false
			}
		}
		return flag
	}else {
		return s.Size()==0
	}
}


/**
删除,成功为true,否则为false
 */
func (s *set) Remove(v interface{}) bool {
	m.Lock()
	flag:=false
	for index,value :=range s.arr  {
		if(reflect.DeepEqual(value, v)){
			flag=true
			s.arr[index]=nil
		}
	}
	defer m.Unlock()
	return flag
}


/**
清除所有的元素
 */
func (s *set) Clear(){
	m.Lock()
	copy(s.arr,make([]interface{},len(s.arr)))
	defer m.Unlock()
}

/**
set变成一个数组进行返回
 */
func (s *set)ToArray() (*[]interface{}) {
	return  &s.arr
}



func main() {
	s:=NewSet()
	fmt.Println(s.IsEmpty())

	/*s.Add(1)
	s.Add("2")
	s.Add(1)
	fmt.Println(s)

	b:=s.Remove(2)
	fmt.Println(b)
	fmt.Println(s)*/

	s.AddAll("a","b","c")
	fmt.Println(s.IsEmpty())
	fmt.Println(s)

	arr:=s.ToArray()
	fmt.Println(arr)

	s.Clear()
	fmt.Println(s.IsEmpty())


}


