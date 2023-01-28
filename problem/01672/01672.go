package _01672

// 1672. Richest Customer Wealth
// https://leetcode.com/problems/richest-customer-wealth/

func maximumWealth(accounts [][]int) int {
	ans := 0
	for _, account := range accounts {
		sum := 0
		for _, price := range account {
			sum += price
		}
		ans = maxInt(ans, sum)
	}
	return ans
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
