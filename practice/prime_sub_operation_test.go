package main

import "testing"

func TestPrimeSubOperation(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Example test",
			nums:     []int{998, 2},
			expected: true,
		},
		{
			name:     "Example 1",
			nums:     []int{4, 9, 6, 10},
			expected: true,
		},
		{
			name:     "Example 2",
			nums:     []int{6, 8, 11, 12},
			expected: true,
		},
		{
			name:     "Example 3",
			nums:     []int{5, 8, 3},
			expected: false,
		},
		{
			name:     "Already sorted",
			nums:     []int{1, 2, 3, 4},
			expected: true,
		},
		{
			name:     "Impossible case",
			nums:     []int{10, 1, 2},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := primeSubOperation(tt.nums)
			if result != tt.expected {
				t.Errorf("primeSubOperation(%v) = %v, expected %v", tt.nums, result, tt.expected)
			}
		})
	}
}
