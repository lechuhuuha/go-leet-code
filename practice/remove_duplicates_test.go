package main

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
		output   []int
	}{
		{
			name:     "Example 1",
			input:    []int{1, 1, 2},
			expected: 2,
			output:   []int{1, 2},
		},
		{
			name:     "Example 2",
			input:    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: 5,
			output:   []int{0, 1, 2, 3, 4},
		},
		{
			name:     "Empty Array",
			input:    []int{},
			expected: 0,
			output:   []int{},
		},
		{
			name:     "Single Element",
			input:    []int{42},
			expected: 1,
			output:   []int{42},
		},
		{
			name:     "No Duplicates",
			input:    []int{1, 2, 3},
			expected: 3,
			output:   []int{1, 2, 3},
		},
		{
			name:     "All Duplicates",
			input:    []int{9, 9, 9, 9},
			expected: 1,
			output:   []int{9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, len(tt.input))
			copy(nums, tt.input)

			// Call function
			result := removeDuplicates(nums)

			// Check the returned count
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}

			// Check the array content up to the returned count
			if !reflect.DeepEqual(nums[:result], tt.output) {
				t.Errorf("expected array %v, got %v", tt.output, nums[:result])
			}
		})
	}
}
