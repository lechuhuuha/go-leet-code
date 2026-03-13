package main

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected []string
	}{
		{
			name:     "n=3",
			n:        3,
			expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name:     "n=1",
			n:        1,
			expected: []string{"()"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := generateParenthesis(tt.n)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("generateParenthesis(%d) = %v, expected %v", tt.n, actual, tt.expected)
			}
		})
	}
}
