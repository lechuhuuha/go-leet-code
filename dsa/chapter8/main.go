package main

import (
	"fmt"

	sorting "github.com/lechuhuuha/go-leet-code/dsa/chapter8/sort"
)

func main() {
	intergers := [11]int{31, 13, 12, 4, 18, 16, 7, 2, 3, 0, 10}
	sorting.BubbleSort(intergers[:])

	elements := []int{11, 4, 18, 6, 19, 21, 71, 13, 15, 2}
	fmt.Println("Before Sorting ", elements)
	sorting.SelectionSort(elements)
	fmt.Println("After Sorting", elements)

	sequence := sorting.RandomSequence(24)
	fmt.Println("\n^^^^^^ Before Sorting ^^^ \n\n", sequence)
	sorting.InsertionSort(sequence)
	fmt.Println("\n--- After Sorting ---\n\n", sequence)

	elements = []int{34, 202, 13, 19, 6, 5, 1, 43, 506, 12, 20, 28, 17, 100,
		25, 4, 5, 97, 1000, 27}
	sorting.ShellSort(elements)
	fmt.Println(elements)
}
