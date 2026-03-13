package main

func booleanMatrix(matrix [3][3]int) [3][3]int {
	var (
		rows          [3]int
		columns       [3]int
		booleanMatrix [3][3]int
	)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if matrix[i][j] == 1 {
				rows[i] = 1
				columns[j] = 1
			}
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if rows[i] == 1 || columns[j] == 1 {
				booleanMatrix[i][j] = 1
			}
		}
	}
	return booleanMatrix
}
