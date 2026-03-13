package main

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"{", false},
		{"{{}", false},
		{"", true},
		{"((()))", true},
		{"((())", false},
		{"))((", false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("isValid(%s)", test.input), func(t *testing.T) {
			result := isValid(test.input)
			if result != test.expected {
				t.Errorf("Expected isValid(%s) to be %v, but got %v", test.input, test.expected, result)
			}
		})
	}
}
