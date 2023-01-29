package _00122

// 122. Best Time to Buy and Sell Stock II
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/

func maxProfit(prices []int) int {
	m := []int{}
	for i := 1; i < len(prices); i++ {
		m = append(m, max(prices[i]-prices[0], 0))
	}

	for i := 1; i < len(prices)-1; i++ {
		tmp := []int{}
		for j := i + 1; j < len(prices); j++ {
			tmp = append(tmp, max(m[j-i], m[0]+max(0, prices[j]-prices[i])))
		}
		m = tmp
	}
	return m[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
