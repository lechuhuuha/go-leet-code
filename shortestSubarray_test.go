package main

import (
	"fmt"
	"testing"
)

func TestShortestSubarray(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			arr:      []int{1},
			k:        1,
			expected: 1,
		},
		{
			name:     "Example 2",
			arr:      []int{1, 2},
			k:        4,
			expected: -1,
		},
		{
			name:     "Example 3",
			arr:      []int{2, -1, 2},
			k:        3,
			expected: 3,
		},
		{
			name:     "Edge Case: Empty Array",
			arr:      []int{},
			k:        1,
			expected: -1,
		},
		{
			name:     "Edge Case: Already Sorted",
			arr:      []int{1, 2, 3, 4},
			k:        10,
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := shortestSubarray(tt.arr, tt.k)
			fmt.Println(tt.arr, tt.k, result)

			if result != tt.expected {
				t.Errorf("shortestSubarray(%v, %d) = %d; want %d", tt.arr, tt.k, result, tt.expected)
			}
		})
	}
}
