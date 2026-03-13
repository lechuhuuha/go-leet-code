package main

func maxDistance(colors []int) int {
	// still the same logic
	// but should save a map to store which colors at which position
	// if the color is the same then move to the next one
	n := len(colors)
	maxDis := 0

	for distance := n - 1; distance >= 0; distance-- {
		if colors[distance] != colors[0] {
			maxDis = max(maxDis, distance)
		}
	}

	for distance := 0; distance < n; distance++ {
		if colors[distance] != colors[n-1] {
			maxDis = max(maxDis, n-1-distance)
		}
	}

	return maxDis
}
