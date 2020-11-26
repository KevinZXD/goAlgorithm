package main

import "container/list"

func link2List(array []int) *list.List {
	listLink := list.New()
	for i := 0; i < len(array); i++ {
		listLink.PushFront(array[i])
	}
	return listLink

}
func main() {
	array := []int{1, 2, 3, 4}
	link := link2List(array)
	for i := link.Front(); i != nil; i = i.Next() {
		print(i.Value)
	}

}
