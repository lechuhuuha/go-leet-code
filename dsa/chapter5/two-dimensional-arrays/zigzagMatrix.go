package main

import "fmt"

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
