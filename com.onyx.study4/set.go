package main


import (
	"fmt"
	"reflect"
)


type Set struct {
	m map[interface{}]int
}

func NewSet() *Set {
	return &Set{
		make(map[interface{}]int),
	}
}

func (s *Set) Add(v interface{}) {
	s.m[v]=1
}

func (s *Set) Remove(v interface{}) {
	delete(s.m,v)
}

func (s *Set) Contains(v interface{}) bool{
	for _,value:= range s.m {
		if(value==v){
			reflect.DeepEqual(value, v)
		}
	}
	return false
}


func (s *Set) Size() int {
	return len(s.m)
}

func (s *Set) Clear() {
	s.m = make(map[interface{}]int)
}

func (s *Set) IsEmpty() bool {
	return len(s.m)==0
}


func main() {
	//初始化
	s := NewSet()
	fmt.Println(s)

	s.Add(1)
	fmt.Println(s)
	s.Add(1)
	fmt.Println(s)

	s.Remove(1)
	fmt.Println(s)

	s.Add(2)
	fmt.Print(s.Size())




}






