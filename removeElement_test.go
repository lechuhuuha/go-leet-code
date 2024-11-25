package main

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name           string
		nums           []int
		val            int
		expectedK      int
		expectedPrefix []int
	}{
		{
			name:           "Example 1",
			nums:           []int{3, 2, 2, 3},
			val:            3,
			expectedK:      2,
			expectedPrefix: []int{2, 2},
		},
		{
			name:           "Example 2",
			nums:           []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:            2,
			expectedK:      5,
			expectedPrefix: []int{0, 1, 4, 0, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums) // Copy nums to avoid modifying the original slice in the test case

			gotK := removeElement(numsCopy, tt.val)

			// Check if the length matches
			if gotK != tt.expectedK {
				t.Errorf("got k = %d, want %d", gotK, tt.expectedK)
			}

			// Check if the prefix matches
			if !reflect.DeepEqual(numsCopy[:gotK], tt.expectedPrefix) {
				t.Errorf("got nums prefix = %v, want %v", numsCopy[:gotK], tt.expectedPrefix)
			}
		})
	}
}
