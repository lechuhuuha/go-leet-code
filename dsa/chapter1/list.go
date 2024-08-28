package main

import (
	"container/heap"
	"container/list"
	"fmt"
)

func main() {
	var intList list.List
	intList.PushBack(11)
	intList.PushBack(23)
	intList.PushBack(34)
	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}

	Tuples()

	var intHeap *IntegerHeap = &IntegerHeap{1, 5, 4}
	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	fmt.Printf("minium: %d\n", (*intHeap)[0])
	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(intHeap))
	}
	fmt.Printf("%v", intHeap)
}
