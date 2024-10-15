package main

import (
	"container/ring"
	"fmt"
)

func main() {
	intergers := []int{1, 3, 5, 7}
	circularList := ring.New(len(intergers))
	for i := 0; i < circularList.Len(); i++ {
		circularList.Value = intergers[i]
		circularList = circularList.Next()
	}

	circularList.Do(func(element interface{}) {
		fmt.Print(element, ",")
	})
	fmt.Println()

	// reverse of the circular list
	for i := 0; i < circularList.Len(); i++ {
		fmt.Print(circularList.Value, ",")
		circularList = circularList.Prev()
	}
	fmt.Println()

	// move two elements forward in the circular list
	circularList = circularList.Move(2)
	circularList.Do(func(element interface{}) {
		fmt.Print(element, ",")
	})
	fmt.Println()
}
