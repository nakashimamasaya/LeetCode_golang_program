package _0136

import "testing"

func TestSample(t *testing.T) {
	nums := [][]int{
		{2, 2, 1},
		{4, 1, 2, 1, 2},
		{1},
	}

	ans := []int{
		1,
		4,
		1,
	}

	for i := 0; i < len(nums); i++ {
		result := SingleNumber(nums[i])
		if result != ans[i] {
			t.Errorf("miss sample %d\n return %d\n answer %d", i+1, result, ans[i])
		}
	}
}
