package main

import (
	"testing"
)

func TestCountFairPairs(t *testing.T) {
	tests := []struct {
		nums     []int
		lower    int
		upper    int
		expected int64
	}{
		{
			nums:     []int{0, 1, 7, 4, 4, 5},
			lower:    3,
			upper:    6,
			expected: 6,
		},
		{
			nums:     []int{1, 7, 9, 2, 5},
			lower:    11,
			upper:    11,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := countFairPairs(tt.nums, tt.lower, tt.upper)
			if result != tt.expected {
				t.Errorf("countFairPairs(%v, %d, %d) = %d; want %d", tt.nums, tt.lower, tt.upper, result, tt.expected)
			}
		})
	}
}
