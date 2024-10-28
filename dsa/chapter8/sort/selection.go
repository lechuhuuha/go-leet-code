package sorting

func SelectionSort(elements []int) {
	for i := 0; i < len(elements)-1; i++ {
		min := i
		for j := i + 1; j < len(elements); j++ {
			if elements[j] < elements[min] {
				min = j
			}
		}
		swap(elements, i, min)
	}
}

func swap(elements []int, i, j int) {
	temp := elements[j]
	elements[j] = elements[i]
	elements[i] = temp
}
