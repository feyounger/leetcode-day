package subject_121

func maxProfit(prices []int) int {
	var max, tmp int
	tmp = prices[0]
	if len(prices) == 1 {
		return 0
	}
	for i := 1; i < len(prices); i++ {
		if prices[i] > tmp {
			if prices[i]-tmp > max {
				max = prices[i] - tmp
			}
		} else {
			tmp = prices[i]
		}
	}
	return max
}
