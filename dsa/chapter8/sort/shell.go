package sorting

func ShellSort(elements []int) {
	intervals := []int{1}
	k := 1
	n := len(elements)
	for {
		interval := Power(2, k) + 1
		if interval > n-1 {
			break
		}
		intervals = append(intervals, interval)
		k++
	}

	for _, interval := range intervals {
		for i := interval; i < n; i += interval {
			j := i
			for j > 0 {
				if elements[j-interval] > elements[j] {
					elements[j-interval], elements[j] = elements[j], elements[j-interval]
				}
				j = j - interval
			}
		}
	}
}

func Power(exponent int, index int) int {
	power := 1
	for index > 0 {
		if index&1 != 0 {
			power *= exponent
		}
		index >>= 1
		exponent *= exponent
	}
	return power
}
