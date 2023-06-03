package _02706

import "sort"

// 2706. Buy Two Chocolates
// https://leetcode.com/problems/buy-two-chocolates/

func buyChoco(prices []int, money int) int {
	sort.Ints(prices)
	if prices[0]+prices[1] > money {
		return money
	}

	return money - prices[0] - prices[1]
}
