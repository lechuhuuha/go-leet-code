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
