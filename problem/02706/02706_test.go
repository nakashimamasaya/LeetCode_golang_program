package _02706

import "testing"

func TestSample(t *testing.T) {
	prices := [][]int{
		{1, 2, 2},
		{3, 2, 3},
	}

	money := []int{
		3,
		3,
	}

	ans := []int{
		0,
		3,
	}

	for i, a := range ans {
		if result := buyChoco(prices[i], money[i]); a != result {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
