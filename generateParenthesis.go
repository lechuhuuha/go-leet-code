package main

func generateParenthesis(n int) []string {
	// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
	// Input: n = 3
	// Output: ["((()))","(()())","(())()","()(())","()()()"]

	result := []string{}
	stack := map[rune]rune{}

	for i := 0; i < n; i++ {
		stack['('] = ')'
	}

	// create the most parenthesis
	// after that minus one
}
