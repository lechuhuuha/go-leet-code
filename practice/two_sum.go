package main

// Given a sorted array A (sorted in asc order), having N integers,
// find if there exists any pair of elements (A[i], A[j]), such that their sum is equal to X

// Start the sum of extreme values (smallest and largest) and conditionally move both pointers.
// Move left pointer "i" by one when the sum of A[i] and A[j] is less then X.
// Move right pointer "j" minus one when the sum of A[i] and A[j] is larger than Y
// We do not miss any pair because the sum is already smaller then X
func twoSum(numbers []int, target int) []int {
	result := []int{}
	first, last := 0, len(numbers)-1
	for first < last {
		rTarget := numbers[first] + numbers[last]
		if rTarget == target {
			return append(result, first+1, last+1)
		}

		if rTarget < target {
			first++
			continue
		}

		if rTarget > target {
			last--
			continue
		}
	}
	return result
}
