package _0053

import "testing"

func TestSample(t *testing.T) {
	nums := [][]int{
		{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		{1},
		{5, 4, -1, 7, 8},
	}

	ans := []int{
		6,
		1,
		23,
	}

	for i := 0; i < len(nums); i++ {
		if MaxSubArray(nums[i]) != ans[i] {
			t.Errorf("miss sample %d", i+1)
		}
	}
}
