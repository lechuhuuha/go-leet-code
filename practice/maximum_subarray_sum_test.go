package main

import "testing"

func TestMaximumSubarraySum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int64
	}{
		{
			name:     "example_1",
			nums:     []int{1, 5, 4, 2, 9, 9, 9},
			k:        3,
			expected: 15,
		},
		{
			name:     "example_2",
			nums:     []int{4, 4, 4},
			k:        3,
			expected: 0,
		},
		{
			name:     "empty_array",
			nums:     []int{},
			k:        3,
			expected: 0,
		},
		{
			name:     "k_greater_than_length",
			nums:     []int{1, 2},
			k:        3,
			expected: 0,
		},
		{
			name:     "no_repeated_elements",
			nums:     []int{1, 2, 3, 4, 5},
			k:        2,
			expected: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximumSubarraySum(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("maximumSubarraySum(%v, %d) = %d; want %d", tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}
