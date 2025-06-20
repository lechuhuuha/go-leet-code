package main

import (
	"reflect"
	"testing"
)

func TestPartitionArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Test case 1",
			nums:     []int{3, 6, 1, 2, 5},
			k:        2,
			expected: 2,
		},
		{
			name:     "Test case 2",
			nums:     []int{1, 2, 3},
			k:        1,
			expected: 2,
		},
		{
			name:     "Test case 3",
			nums:     []int{2, 2, 4, 5},
			k:        0,
			expected: 3,
		},
		{
			name:     "Test case 4",
			nums:     []int{0, 0, 0, 0, 0},
			k:        0,
			expected: 1,
		},
		{
			name:     "Test case 5",
			nums:     []int{3, 1, 3, 4, 2},
			k:        0,
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := partitionArray(tt.nums, tt.k)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("generate(%d) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
