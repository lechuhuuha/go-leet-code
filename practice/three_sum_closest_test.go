package main

import (
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{-1, 2, 1, -4}, 1, 2},
		{[]int{0, 0, 0}, 1, 0},
		{[]int{1, 1, 1, 1}, 2, 3},
		{[]int{1, 1, -1, -1, 3}, -1, -1},
		{[]int{-1, 0, 1, 1}, 1, 1},
		{[]int{-1, 2, 1, -4, 3}, 1, 1},
	}

	for _, test := range tests {
		result := threeSumClosest(test.nums, test.target)
		if result != test.expected {
			t.Errorf("For nums %v and target %d, expected %d, but got %d", test.nums, test.target, test.expected, result)
		}
	}
}
