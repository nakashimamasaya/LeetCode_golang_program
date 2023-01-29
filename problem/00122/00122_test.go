package _00122

import "testing"

func TestSample(t *testing.T) {
	prices := [][]int{
		{7, 1, 5, 3, 6, 4},
		{1, 2, 3, 4, 5},
		{7, 6, 4, 3, 1},
	}

	ans := []int{
		7,
		4,
		0,
	}

	for i, p := range prices {
		if result := maxProfit(p); ans[i] != result {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
