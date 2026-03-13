package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			fmt.Println(stack)
			// Pop the last two numbers from the stack
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			// Perform the operation
			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				stack = append(stack, a/b)
			}
		default:
			// Convert the token to an integer and push it onto the stack
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}
