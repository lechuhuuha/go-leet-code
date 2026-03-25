package problems

// You are given two integer arrays nums1 and nums2 where:
// nums1 is a subset of nums2
// for each element in nums1, find the next greater element in nums2
// The next greater element of x in nums2 is:
// the first element to the right of x
// that is greater than x
// if none exists, answer is -1

// nums1 = [4,1,2]
// nums2 = [1,3,4,2]

// answer = [-1,3,-1]
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	output := make([]int, len(nums1))
	m := make(map[int]int, len(nums2))
	stack := monotonicStack{}
	for i := range nums1 {
		output[i] = -1
	}

	for i := 0; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[i] > stack.Peek() {
			val := stack.Pop()
			m[val] = nums2[i]
		}
		stack.Push(nums2[i])
	}

	for i, v := range nums1 {
		if val, ok := m[v]; ok {
			output[i] = val
		}
	}
	return output
}
