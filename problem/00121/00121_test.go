package _00121

import "testing"

func TestSample(t *testing.T) {
	prices := [][]int{
		{7, 1, 5, 3, 6, 4},
		{7, 6, 4, 3, 1},
	}

	ans := []int{
		5, 0,
	}

	for i := 0; i < len(prices); i++ {
		result := maxProfit(prices[i])
		if result != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
