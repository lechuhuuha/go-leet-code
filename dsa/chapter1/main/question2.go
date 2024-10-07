package main

func MaxMinValue(values []int) (max int, min int) {
	if len(values) == 0 {
		return 0, 0
	}
	min, max = values[0], values[0]
	for _, v := range values {
		if max < v {
			max = v
		}

		if min > v {
			min = v
		}
	}

	return
}
