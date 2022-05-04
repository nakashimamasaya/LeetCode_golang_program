package _0704

import "testing"

func TestSample(t *testing.T) {
	nums := [][]int{
		{-1, 0, 3, 5, 9, 12},
		{-1, 0, 3, 5, 9, 12},
	}

	target := []int{
		9,
		2,
	}

	ans := []int{
		4,
		-1,
	}

	for i := 0; i < len(nums); i++ {
		result := Search(nums[i], target[i])
		if ans[i] != result {
			t.Errorf("miss sample %d\nreturn %v\nanswer %v", i+1, result, ans[i])
		}
	}
}
