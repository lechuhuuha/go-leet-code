package main

import (
	"reflect"
	"testing"
)

func TestGetRow(t *testing.T) {
	tests := []struct {
		name     string
		rowIndex int
		expected []int
	}{
		{
			name:     "Row 0",
			rowIndex: 0,
			expected: []int{1},
		},
		{
			name:     "Row 1",
			rowIndex: 1,
			expected: []int{1, 1},
		},
		{
			name:     "Row 2",
			rowIndex: 2,
			expected: []int{1, 2, 1},
		},
		{
			name:     "Row 3",
			rowIndex: 3,
			expected: []int{1, 3, 3, 1},
		},
		{
			name:     "Row 4",
			rowIndex: 4,
			expected: []int{1, 4, 6, 4, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pascalTriangleTwo(tt.rowIndex)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("getRow(%d) = %v, want %v", tt.rowIndex, result, tt.expected)
			}
		})
	}
}
