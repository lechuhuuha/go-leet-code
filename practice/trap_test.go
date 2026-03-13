package main

import (
	"testing"
)

func TestTrap(t *testing.T) {
	tests := []struct {
		height   []int
		expected int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
		{[]int{}, 0},
		{[]int{2, 0, 2}, 2},
		{[]int{3, 1, 2, 1, 3}, 5},
		{[]int{1, 2, 3, 4, 5}, 0},
		{[]int{5, 4, 3, 2, 1}, 0},
	}

	for _, test := range tests {
		result := trap(test.height)
		if result != test.expected {
			t.Errorf("For height %v, expected %d, but got %d", test.height, test.expected, result)
		}
	}
}
