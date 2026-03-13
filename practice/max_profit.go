package main

func maxProfit(prices []int) int {
	minPrices := prices[0]
	maxProfit := 0

	for day := 0; day < len(prices); day++ {
		if prices[day] > minPrices {
			maxProfit = max((prices[day] - minPrices), maxProfit)
		} else {
			minPrices = prices[day]
		}
	}
	return maxProfit
}
