package _01672

import "testing"

func TestSample(t *testing.T) {
	accounts := [][][]int{
		{{1, 2, 3}, {3, 2, 1}},
		{{1, 5}, {7, 3}, {3, 5}},
		{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}},
	}

	ans := []int{
		6,
		10,
		17,
	}

	for i, account := range accounts {
		if result := maximumWealth(account); result != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
