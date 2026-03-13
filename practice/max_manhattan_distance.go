package main

import "math"

func maxManhattanDistance(str string, k int) int {
	/*
	   changing 'N' to 'S' => dy -= 2
	   N to E => dx += 1, dy -= 1 => manhattan no change
	   N to W => dx -= 1, dy -= 1 > manhattan no change
	   N to N => no change

	   similarly, for other directions as well
	   we need to now maximize by flipping (N <=> S) and (E <=> W) atmax k

	   iterate and at each pos, keep track of news count we have seen
	   max distance, that can be obtained is replacing (N/S) and (E/W) with k replacements
	   choose minfreq of N, S and replace all those with min(k, minfreq),
	   with remaining (k - minfreq), we could replace (E/W)

	   in essence, min(k, (minFreq(N/S) + minFreq(E/W))) replacements are chosen everytime

	   find the max of all these values while iterating.
	*/

	ans := 0
	n, e, w, s := 0, 0, 0, 0
	for _, val := range str {
		if val == 'N' {
			n++
		}
		if val == 'E' {
			e++
		}
		if val == 'W' {
			w++
		}
		if val == 'S' {
			s++
		}

		distChange := 2 * min(k, min(n, s)+min(e, w))
		distWithoutChanges := int(math.Abs(float64(n-s))) + int(math.Abs(float64(e-w)))

		ans = max(ans, distWithoutChanges+distChange)
	}
	return ans
}
