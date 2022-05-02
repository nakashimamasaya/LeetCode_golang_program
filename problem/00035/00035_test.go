package _0035

import "testing"

func TestSample(t *testing.T) {
	nums := [][]int{
		{1, 3, 5, 6},
		{1, 3, 5, 6},
	}
	target := []int{5, 2}
	ans := []int{2, 1}

	for i := 0; i < 2; i++ {
		if SearchInsert(nums[i], target[i]) != ans[i] {
			t.Errorf("miss sample %d", i+1)
		}
	}
}
