package _00121

import "math"

// 121. Best Time to Buy and Sell Stock
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

func maxProfit(prices []int) int {
	priceLen := len(prices)
	memo := make([]float64, priceLen)
	memo[priceLen-1] = float64(prices[priceLen-1])
	for i := priceLen - 2; i >= 0; i-- {
		memo[i] = math.Max(
			float64(memo[i+1]),
			float64(prices[i]),
		)
	}
	ans := float64(0)
	for i := 0; i < priceLen; i++ {
		ans = math.Max(
			ans,
			memo[i]-float64(prices[i]),
		)
	}
	return int(ans)
}
