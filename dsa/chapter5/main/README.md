1. What is 2-mode unfolding of a tensor array?
   The 2-mode unfolding of a tensor array is shown here. The array's first dimension row
   index is set to 2:

```go
   for j=0; j < 3; j++ {
    for k=0; k < 3; k++ {
    fmt.Printf("%d ",array[2][j][k])
    }
   fmt.Printf("\n")
   }
```

2. Write a two-dimensional array of strings and initialize it. Print the strings.

```go
	twoDimensionalArr := [2][2]int{
		{2, 3},
		{23, 3},
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(twoDimensionalArr[i][j])
		}
	}
```

3. Give an example of a multi-dimensional array and traverse through it.

```go
	threedimensionalArr := [2][2][2]int{}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				threedimensionalArr[i][j][k] = rand.Intn(3)
			}
		}
	}
```

4. For a 3 x 3 matrix, write code that calculates the determinant of the matrix.

det=a(ei−fh)−b(di−fg)+c(dh−eg)
```go
func determinant3x3(matrix [3][3]int) (def float32) {
	def = float32(
		matrix[0][0]*(matrix[1][1]*matrix[2][2]-matrix[1][2]*matrix[2][1]) -
			matrix[0][1]*(matrix[1][0]*matrix[2][2]-matrix[1][2]*matrix[2][0]) +
			matrix[0][2]*(matrix[1][0]*matrix[2][1]-matrix[1][1]*matrix[2][0]))
	return
}
```

5. What is a transpose of a 3 x 3 matrix?

```go
func transpose3x3(matrix [3][3]int) [3][3]int {
	var transMatrix [3][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			transMatrix[i][j] = matrix[j][i]
		}
	}

	return transMatrix
}
```

6. What is a zig-zag matrix?

- A zig-zag matrix is a square arrangement of n x n integers. The integers are arranged on
  anti-diagonals in sequentially increasing order.

```go
func createZigzagMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	// Variables to track the current position
	row, col := 0, 0
	goingDown := true // Zigzag direction

	for num := 1; num <= n*n; num++ {
		matrix[row][col] = num
		fmt.Printf("Placing %d at (%d, %d)\n", num, row, col) // Log the placement

		if goingDown {
			if col == 0 || row == n-1 { // If it hits the left or bottom edge
				goingDown = false
				if row == n-1 { // If at bottom, move right
					col++
				} else { // If at left, move down
					row++
				}
			} else { // Continue moving diagonally down
				row++
				col--
			}
		} else {
			if row == 0 || col == n-1 { // If it hits the top or right edge
				goingDown = true
				if col == n-1 { // If at right, move down
					row++
				} else { // If at top, move right
					col++
				}
			} else { // Continue moving diagonally up
				row--
				col++
			}
		}
	}

	return matrix
}
```

7. Write code with an example of a spiral matrix.

```go
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
```

8. Which dimension is typically unfolded for tensor arrays?

- three dimensional array and its unfold 1 mode

9. How do you define a Boolean matrix?

- A Boolean matrix is a matrix that consists of elements in the mth row and the nth column with
  a value of 1.
- A matrix can be modified to be a Boolean matrix by making the values in
  the mth row and the nth column equal to 1.

10. Choose two 3 x 3 matrices and find the product of the matrices.

```go
for l := 0; l < 3; l++ { // Outer loop for rows of matrix1
		for m := 0; m < 3; m++ { // Inner loop for columns of matrix2
			productSum := 0
			for n := 0; n < 3; n++ { // Loop to multiply rows of matrix1 with columns of matrix2
				productSum += matrix1[l][n] * matrix2[n][m]
			}
			product[l][m] = productSum // Store the computed product
		}
	}
```
