package main

import (
	"reflect"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	tests := []struct {
		name     string
		input    []any
		expected []int
	}{
		{
			name:     "simple right-skewed tree",
			input:    []any{1, nil, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "complex tree",
			input:    []any{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9},
			expected: []int{1, 2, 4, 5, 6, 7, 3, 8, 9},
		},
		{
			name:     "empty tree",
			input:    []any{},
			expected: []int{},
		},
		{
			name:     "single node",
			input:    []any{1},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.input)
			got := preorderTraversal(root)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("preorderTraversal() = %v, want %v", got, tt.expected)
			}
		})
	}
}
