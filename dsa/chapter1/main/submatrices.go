package main

import "fmt"

// Function to extract and print 2x2 submatrices
func find2x2Submatrices(matrix [][]int) {
	// Check if the matrix is at least 3x3
	if len(matrix) < 3 || len(matrix[0]) < 3 {
		fmt.Println("Matrix is too small to extract 2x2 submatrices")
		return
	}

	fmt.Println("2x2 Submatrices:")
	// Loop through all possible 2x2 submatrices
	// go from top to bottom
	for i := 0; i < len(matrix)-1; i++ {
		for j := 0; j < len(matrix[i])-1; j++ {
			// Extract the 2x2 submatrix
			submatrix := [][]int{
				{matrix[i][j], matrix[i][j+1]},
				{matrix[i+1][j], matrix[i+1][j+1]},
			}
			// Print the submatrix
			fmt.Println(submatrix)
		}
	}
}
