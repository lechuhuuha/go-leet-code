package main

import (
	"container/heap"
	"fmt"
)

func main() {
	// tuples
	Tuples()

	// heap
	var intHeap *IntegerHeap = &IntegerHeap{1, 5, 4}
	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	fmt.Printf("minium: %d\n", (*intHeap)[0])
	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(intHeap))
	}
	fmt.Printf("%v", intHeap)

	// adapter
	var processor IProcess = Adapter{}
	processor.process()

	// bridge
	var x = [5]float32{1, 2, 3, 4, 5}
	var y = [5]float32{1, 2, 3, 4, 5}
	var contour IContour = DrawContour{
		x:      x,
		y:      y,
		shape:  DrawShape{},
		factor: 2,
	}
	contour.drawContour(x, y)
	contour.resizeByFactor(2)

	// composite
	var mainBranch = &Branch{name: "Branch 1"}
	var mainLeaf = Leaflet{name: "Leaf 1"}
	var subBranch = &Branch{name: "Branch 2"}
	var subLeaf = Leaflet{name: "Leaf 2"}

	mainBranch.add(mainLeaf)
	mainBranch.add(subLeaf)
	mainBranch.addBranch(*subBranch)

	mainBranch.perform()

	// decorator
	var processClass = ProcessClass{}
	var decorator = ProcessDecorator{}
	decorator.process()
	decorator.processInstance = &processClass
	decorator.process()
}
