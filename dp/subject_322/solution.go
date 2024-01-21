package subject_322

import "math"

func coinChange(coins []int, amount int) int {
	m := make(map[int]int, amount+1)
	var coin func([]int, int) int
	coin = func(coins []int, amount int) int {
		var minFun = func(a, b int) int {
			if a < b {
				return a
			}
			return b
		}
		if amount == 0 {
			return 0
		}
		if amount < 0 {
			return -1
		}
		if value, ok := m[amount]; ok {
			return value
		}
		res := math.MaxInt32
		for _, money := range coins {
			tmp := coin(coins, amount-money)
			if tmp == -1 {
				continue
			}
			res = minFun(res, tmp+1)
		}
		if res == math.MaxInt32 {
			m[amount] = -1
		} else {
			m[amount] = res
		}
		return m[amount]
	}

	return coin(coins, amount)
}
