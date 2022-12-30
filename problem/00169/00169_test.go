package _00169

import "testing"

func TestSample(t *testing.T) {
	nums := [][]int{
		{3, 2, 3},
		{2, 2, 1, 1, 1, 2, 2},
	}

	ans := []int{
		3,
		2,
	}

	for i, n := range nums {
		if result := majorityElement(n); result != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
