package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	// mỗi dòng sẽ có từ 1-9 mà ko lặp lại
	// mỗi cột sẽ có từ 1-9 mà ko lặp lại
	// mỗi subbox (gồm 3 dòng và 3 cột ) sẽ có từ 1-9 mà ko được lặp lại

	isValid := true
	rows := [9][9]bool{}
	cols := [9][9]bool{}
	boxes := [9][9]bool{}
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				continue
			}
			num := board[row][col] - '1'

			// check row
			if rows[row][num] {
				return false
			}

			rows[row][num] = true

			// check col

			if cols[col][num] {
				return false
			}
			cols[col][num] = true

			// check box
			boxIndex := (row/3)*3 + col/3
			fmt.Printf("row : %d, col : %d, boxindex : %d, num : %d\n", row, col, boxIndex, num)
			if boxes[boxIndex][num] {
				return false
			}
			boxes[boxIndex][num] = true
		}
	}
	return isValid
}

func isValidRow(board [][]byte, row int) bool {
	var nums [10]bool
	for col := 0; col < 9; col++ {
		n := convertToNumber(board[row][col])
		if n < 0 {
			continue
		}
		// lần 1 thì giá trị là false nên lần 2 là coi như lặp lại trên cùng 1 row nên bị coi là false
		if nums[n] {
			return false
		}
		nums[n] = true
	}
	return true
}

func isValidCol(board [][]byte, col int) bool {
	var nums [10]bool
	for row := 0; row < 9; row++ {
		n := convertToNumber(board[row][col])
		if n < 0 {
			continue
		}
		if nums[n] {
			return false
		}
		nums[n] = true
	}
	return true
}

func isValidPod(board [][]byte, pod int) bool {
	var nums [10]bool

	row := (pod / 3) * 3
	col := (pod % 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			n := convertToNumber(board[row+i][col+j])
			if n < 0 {
				continue
			}
			if nums[n] {
				return false
			}
			nums[n] = true
		}
	}
	return true
}

func convertToNumber(b byte) int {
	if b == '.' {
		return -1
	}
	return int(b - '0')
}
