package main

import "fmt"

func main() {
	Linear()

	var tree *Tree = &Tree{nil, 1, nil}
	print(tree)
	tree.insert(3)
	print(tree)
	tree.insert(5)
	print(tree)
	tree.LeftNode.insert(7)
	print(tree)

	var arr = [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}
	var check bool = findElement(arr, 10)
	fmt.Println(check)
	var check2 bool = findElement(arr, 9)
	fmt.Println(check2)

	arrBinary := []int{1, 2, 3, 6, 8, 9}
	k := 8

	if idx, found := findElementBinarySearch(arrBinary, k); found {
		fmt.Printf("Element %d found at index %d using binary search\n", k, idx)
	} else {
		fmt.Printf("Element %d not found using binary search\n", k)
	}

	var m int
	for m = 0; m < 8; m++ {
		var fib = fibonacci(m)
		fmt.Println(fib)
	}

	arr = [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}
	var addedSum int = 18
	var combinations [19]int
	findElementsWithSum(arr, combinations, 10, addedSum, 0, 0, 0)
}
