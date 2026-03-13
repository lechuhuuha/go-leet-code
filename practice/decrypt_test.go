package main

import (
	"reflect"
	"testing"
)

func TestDecrypt(t *testing.T) {
	tests := []struct {
		name     string
		code     []int
		k        int
		expected []int
	}{
		{
			name:     "Positive k",
			code:     []int{5, 7, 1, 4},
			k:        3,
			expected: []int{12, 10, 16, 13},
		},
		{
			name:     "Zero k",
			code:     []int{1, 2, 3, 4},
			k:        0,
			expected: []int{0, 0, 0, 0},
		},
		{
			name:     "Negative k",
			code:     []int{2, 4, 9, 3},
			k:        -2,
			expected: []int{12, 5, 6, 13},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := decrypt(tt.code, tt.k)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("decrypt(%v, %d) = %v; want %v", tt.code, tt.k, result, tt.expected)
			}
		})
	}
}
