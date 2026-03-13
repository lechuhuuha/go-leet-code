package main

func maxArea(height []int) int {
	maxArea := 0
	if len(height) == 0 {
		return maxArea
	}
	left, right := 0, len(height)-1

	for left < right {
		h := min(height[left], height[right])
		width := right - left
		area := h * width

		if area > maxArea {
			maxArea = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}

	}

	return maxArea
}
