package main

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "Target exists in array",
			nums:     []int{1, 3, 5, 6},
			target:   5,
			expected: 2,
		},
		{
			name:     "Target would be inserted in the middle",
			nums:     []int{1, 3, 5, 6},
			target:   2,
			expected: 1,
		},
		{
			name:     "Target would be inserted at the end",
			nums:     []int{1, 3, 5, 6},
			target:   7,
			expected: 4,
		},
		{
			name:     "Target would be inserted at the beginning",
			nums:     []int{1, 3, 5, 6},
			target:   0,
			expected: 0,
		},
		{
			name:     "Target would be inserted in the middle",
			nums:     []int{1, 3},
			target:   2,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchInsert(tt.nums, tt.target)
			if got != tt.expected {
				t.Errorf("searchInsert(%v, %d) = %d; want %d", tt.nums, tt.target, got, tt.expected)
			}
		})
	}
}
