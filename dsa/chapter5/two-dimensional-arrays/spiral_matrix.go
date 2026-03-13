package main

func PrintSpiral(matrix [][]int) (out []int) {
	n := len(matrix)
	m := len(matrix[0])

	left := 0
	right := m - 1
	top := 0
	bottom := n - 1

	for left <= right && top <= bottom {
		// left move
		for col := left; col <= right; col++ {
			out = append(out, matrix[top][col])
		}
		top++

		// downward move
		for row := top; row <= bottom; row++ {
			out = append(out, matrix[row][right])
		}
		right--

		if left > right || top > bottom {
			break
		}

		// right move
		for col := right; col >= left; col-- {
			out = append(out, matrix[bottom][col])
		}
		bottom--

		// upward move
		for row := bottom; row >= top; row-- {
			out = append(out, matrix[row][left])
		}
		left++
	}
	return
}
