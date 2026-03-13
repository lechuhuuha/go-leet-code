package main

import "testing"

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		tokens   []string
		expected int
	}{
		{tokens: []string{"2", "1", "+", "3", "*"}, expected: 9},
		{tokens: []string{"4", "13", "5", "/", "+"}, expected: 6},
		{tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, expected: 22},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := evalRPN(test.tokens)
			if result != test.expected {
				t.Errorf("For tokens %v, expected %d, but got %d", test.tokens, test.expected, result)
			}
		})
	}
}
