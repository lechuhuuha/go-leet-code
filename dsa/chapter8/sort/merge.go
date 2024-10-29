package sorting

import "fmt"

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	middle := len(arr) / 2

	fmt.Println(arr[:middle], arr[middle:])
	return JoinArrays(MergeSort(arr[:middle]), MergeSort(arr[middle:]))
}

func JoinArrays(leftArr, rightArr []int) []int {
	num, i, j := len(leftArr)+len(rightArr), 0, 0
	arr := make([]int, num)

	for k := 0; k < num; k++ {
		if i > len(leftArr)-1 && j <= len(rightArr)-1 {
			arr[k] = rightArr[j]
			j++
		} else if j > len(rightArr)-1 && i <= len(leftArr)-1 {
			arr[k] = leftArr[i]
			i++
		} else if leftArr[i] < rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
	}
	return arr
}
