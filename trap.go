package main

// [0,1,0,2,1,0,1,3,2,1,2,1]
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	waterTrapped := 0
	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]

	for left < right {
		if height[left] < height[right] {
			left++
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				waterTrapped += leftMax - height[left]
			}
		} else {
			right--
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				waterTrapped += rightMax - height[right]
			}
		}
	}

	return waterTrapped
}
