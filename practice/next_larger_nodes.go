package main

/*
*
  - Definition for singly-linked list.
*/
func nextLargerNodes(head *ListNode) []int {
	values := []int{}
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}

	res := make([]int, len(values))
	stack := []int{} // stack of indices

	for i := 0; i < len(values); i++ {
		for len(stack) > 0 && values[i] > values[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[idx] = values[i]

		}
		stack = append(stack, i)
	}

	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}
