package main

import (
	"container/list"
	"fmt"
	"sync"
)

func gosync() {
	syn := &sync.WaitGroup{}
	for {
		syn.Add(2)
		go func(s *sync.WaitGroup) {
			print(1)
			defer s.Done()
		}(syn)
		go func(s *sync.WaitGroup) {
			print(2)
			defer s.Done()
		}(syn)
		syn.Wait()
	}

}
func gochan() {

	ch := make(chan int, 1)
	ch <- 1
	defer close(ch)
	for {
		go func(ch chan int) {
			<-ch
			print(1)
			ch <- 1

		}(ch)
		go func(ch chan int) {
			<-ch
			print(2)
			ch <- 1
		}(ch)
	}

}
func syncMap() {
	map1 := sync.Map{}
	x, _ := map1.Load("name")
	print(x)
	map1.Store("name", "zxd")
	y, _ := map1.Load("name")
	print(y)
}
func main() {
	l := list.New()
	// 尾部添加
	l.PushBack("canon")
	// 头部添加
	l.PushBack(67)
	l.PushBack(68)
	var next *list.Element
	for i := l.Front(); i != nil; i = next {
		next = i.Next()
		fmt.Println(i.Value)
		l.Remove(i)

	}

	fmt.Println(l.Len())
}
