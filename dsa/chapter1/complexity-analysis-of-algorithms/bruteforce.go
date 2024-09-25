package main

// findElement method given array and k element
func findElement(arr [10]int, k int) bool {
	var i int
	for i = 0; i < 10; i++ {
		if arr[i] == k {
			return true
		}
	}
	return false
}

// Binary search method to find the element k in a sorted array
func findElementBinarySearch(arr []int, k int) (int, bool) {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == k {
			return mid, true
		}
		if arr[mid] < k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1, false
}
