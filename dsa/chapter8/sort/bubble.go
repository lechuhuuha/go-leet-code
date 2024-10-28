package sorting

import "fmt"

func BubbleSort(arr []int) {
	isSwapped := true
	for isSwapped {
		isSwapped = false
		for i := 1; i < len(arr); i++ {
			// if current element is smaller then prev element
			// then move current backward and prev to current element
			if arr[i-1] > arr[i] {
				temp := arr[i]
				arr[i] = arr[i-1]
				arr[i-1] = temp
				isSwapped = true
			}
		}
	}
	fmt.Println(arr)
}
