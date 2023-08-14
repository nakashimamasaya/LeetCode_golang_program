package _02815

import "testing"

func TestSample(t *testing.T) {
	nums := [][]int{
		{51, 71, 17, 24, 42},
		{1, 2, 3, 4},
	}

	ans := []int{
		88,
		-1,
	}

	for i, n := range nums {
		if result := maxSum(n); ans[i] != result {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
