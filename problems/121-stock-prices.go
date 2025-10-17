package problems

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min := prices[0]
	maxProfit := 0

	for _, price := range prices {
		if price < min {
			min = price
			continue
		}
		if price-min > maxProfit {
			maxProfit = price - min
		}
	}
	return maxProfit
}
