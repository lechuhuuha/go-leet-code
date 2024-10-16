package main

// isSymmetric checks if a given matrix is symmetric.
func IsSymmetric(matrix [][]int) bool {
	n := len(matrix)
	// Ensure matrix is square
	for i := 0; i < n; i++ {
		if len(matrix[i]) != n {
			return false
		}
	}
	// Check symmetry
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if matrix[i][j] != matrix[j][i] {
				return false
			}
		}
	}
	return true
}
