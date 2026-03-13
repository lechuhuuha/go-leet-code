package main

import (
	"reflect"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		items    []int
		expected int
	}{
		{
			items:    []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			items:    []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{
			items:    []int{1, 5, 2, 10},
			expected: 9,
		},
	}

	for _, tt := range tests {
		t.Run("maximumBeauty test case", func(t *testing.T) {
			result := maxProfit(tt.items)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
