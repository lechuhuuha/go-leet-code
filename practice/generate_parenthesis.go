package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	var result []string
	var backtrack func(current string, open, close int)

	backtrack = func(current string, open, close int) {
		fmt.Printf("Called backtrack: current='%s', open=%d, close=%d\n", current, open, close)

		// Base case: if the current string length is 2 * n, it's a valid combination
		if len(current) == 2*n {
			fmt.Printf("Adding valid combination: %s\n", current)
			result = append(result, current)
			return
		}

		// If we can still add an open parenthesis, add it
		if open < n {
			fmt.Println("Adding '(': open < n")
			backtrack(current+"(", open+1, close)
		}

		// If we can add a close parenthesis (must not exceed open), add it
		if close < open {
			fmt.Println("Adding ')': close < open")
			backtrack(current+")", open, close+1)
		}

		// When backtracking occurs
		fmt.Printf("Backtracking from: current='%s', open=%d, close=%d\n", current, open, close)
	}

	// Start the backtracking process
	backtrack("", 0, 0)

	return result
}
