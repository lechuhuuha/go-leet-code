package main

import (
	"reflect"
	"testing"
)

func TestMaximumDifference(t *testing.T) {
	tests := []struct {
		items    []int
		expected int
	}{
		{
			items:    []int{7, 1, 5, 4},
			expected: 4,
		},
		{
			items:    []int{9, 4, 3, 2},
			expected: -1,
		},
		{
			items:    []int{1, 5, 2, 10},
			expected: 9,
		},
	}

	for _, tt := range tests {
		t.Run("maximumBeauty test case", func(t *testing.T) {
			result := maximumDifference(tt.items)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
