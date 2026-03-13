package main

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += carry
		if digits[i] < 10 {
			carry = 0
			break
		}
		digits[i] = 0
	}

	if carry > 0 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
