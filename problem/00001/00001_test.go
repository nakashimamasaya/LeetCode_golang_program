package _0001

import "testing"

func TestSample(t *testing.T) {
	nums := [][]int{
		{2, 7, 11, 15},
		{3, 2, 4},
		{3, 3},
	}

	target := []int{
		9,
		6,
		6,
	}

	ans := [][]int{
		{0, 1},
		{1, 2},
		{0, 1},
	}

	for i := 0; i < len(ans); i++ {
		result := twoSum(nums[i], target[i])
		for j := 0; j < len(ans[i]); j++ {
			if result[j] != ans[i][j] {
				t.Errorf("miss sample %v\n return %v \n answer %v", i+1, result, ans[i])
				break
			}
		}

	}
}
