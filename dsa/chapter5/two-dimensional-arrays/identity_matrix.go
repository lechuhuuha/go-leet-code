package main

func IdentityMatrix(order int) [][]float64 {
	identityMatrix := make([][]float64, order)

	for i := 0; i < order; i++ {
		temp := make([]float64, order)
		temp = append(temp, 1)
		identityMatrix = append(identityMatrix, temp)
	}
	return identityMatrix
}
