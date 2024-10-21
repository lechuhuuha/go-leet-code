package main

var matrixOps = [2][2]int{
	{4, 5},
	{1, 2},
}

var maxtrisOpsSecond = [2][2]int{
	{6, 7},
	{3, 4},
}

func addMatrices(matrix1, maxtrix2 [2][2]int, ops string) [2][2]int {
	var sum [2][2]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			result := 0
			if ops == "+" {
				result = matrix1[i][j] + maxtrix2[i][j]
			}
			if ops == "-" {
				result = matrix1[i][j] - maxtrix2[i][j]
			}
			if ops == "*" {
				productSum := 0
				for n := 0; i < 2; i++ {
					productSum = productSum + matrix1[i][j]*maxtrix2[n][j]
				}
				result = productSum
			}
			sum[i][j] = result
		}
	}
	return sum
}

func transpose(matrix [2][2]int) [2][2]int {
	var transMatrix [2][2]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			transMatrix[i][j] = matrix[j][i]
		}
	}

	return transMatrix
}

func transpose3x3(matrix [3][3]int) [3][3]int {
	var transMatrix [3][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			transMatrix[i][j] = matrix[j][i]
		}
	}

	return transMatrix
}

// 1 2 3
// 4 5 6
// 7 8 9
func determinant(matrix [2][2]int) (def float32) {
	def = float32(matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0])
	return
}

func determinant3x3(matrix [3][3]int) (def float32) {
	def = float32(
		matrix[0][0]*(matrix[1][1]*matrix[2][2]-matrix[1][2]*matrix[2][1]) -
			matrix[0][1]*(matrix[1][0]*matrix[2][2]-matrix[1][2]*matrix[2][0]) +
			matrix[0][2]*(matrix[1][0]*matrix[2][1]-matrix[1][1]*matrix[2][0]))
	return
}

func inverse(matrix [2][2]int) [2][2]float64 {
	det := float64(determinant(matrix))
	var inverseMaxtrix [2][2]float64
	inverseMaxtrix[0][0] = float64(matrix[1][1]) / det
	inverseMaxtrix[0][1] = -float64(matrix[0][1]) / det
	inverseMaxtrix[1][0] = -float64(matrix[1][0]) / det
	inverseMaxtrix[1][1] = -float64(matrix[1][1]) / det

	return inverseMaxtrix
}
