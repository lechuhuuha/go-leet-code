package main

import "fmt"

func main() {
	fmt.Println(MaxMinValue([]int{0, 3, 4, 42354, 54354234, 12312321, 432, 1, 12343, 53, 123123, 543543543, 12323, 43543}))

	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	find2x2Submatrices(matrix)
}
