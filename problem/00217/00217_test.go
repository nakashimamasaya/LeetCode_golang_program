package _0217

import "testing"

func TestSample(t *testing.T) {
	nums := [][]int{
		{1, 2, 3, 1},
		{1, 2, 3, 4},
		{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
	}

	ans := []bool{
		true,
		false,
		true,
	}

	for i := 0; i < len(nums); i++ {
		result := ContainsDuplicate(nums[i])
		if result != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
