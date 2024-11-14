package main

import (
	"testing"
)

func TestMinimizedMaximum(t *testing.T) {
	tests := []struct {
		name       string
		n          int
		quantities []int
		expected   int
	}{
		{
			name:       "Example 1",
			n:          6,
			quantities: []int{11, 6},
			expected:   3,
		},
		{
			name:       "Example 2",
			n:          7,
			quantities: []int{15, 10, 10},
			expected:   5,
		},
		{
			name:       "Example 3",
			n:          1,
			quantities: []int{100000},
			expected:   100000,
		},
		{
			name:       "Example 4",
			n:          2,
			quantities: []int{5, 7},
			expected:   7,
		},
		{
			name:       "Example 5",
			n:          3,
			quantities: []int{2, 10, 6},
			expected:   10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minimizedMaximum(tt.n, tt.quantities)
			if result != tt.expected {
				t.Errorf("got %d, expected %d", result, tt.expected)
			}
		})
	}
}
