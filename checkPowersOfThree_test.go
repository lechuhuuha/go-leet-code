package main

import "testing"

func TestCheckPowersOfThree(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{
			name:     "number 12",
			input:    12,
			expected: true,
		},
		{
			name:     "number 91",
			input:    91,
			expected: true,
		},
		{
			name:     "number 21",
			input:    21,
			expected: false,
		},
		{
			name:     "zero input",
			input:    0,
			expected: true,
		},
		{
			name:     "number 45",
			input:    45,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkPowersOfThree(tt.input)
			if got != tt.expected {
				t.Errorf("checkPowersOfThree(%d) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
