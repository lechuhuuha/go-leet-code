package main

import (
	"fmt"
	"strings"
)

// PAYPALISHIRING
// 3
func zigzagconversion(s string, numRows int) string {
	if len(s) <= 1 || len(s) >= 1000 {
		return s
	}
	if numRows <= 1 || numRows >= 1000 {
		return s
	}
	result := make([]string, numRows)
	index, step := 0, 1
	for i := 0; i < len(s); i++ {
		result[index] += string(s[i])
		if index == 0 {
			step = 1
		}
		if index == numRows-1 {
			step = -1
		}
		fmt.Println(index, string(s[i]), step, i, result)

		index += step
	}
	return strings.Join(result, "")
}
