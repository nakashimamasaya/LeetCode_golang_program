package _0219

import "testing"

func TestSample(t *testing.T) {
	nums := [][]int{
		{1, 2, 3, 1},
		{1, 0, 1, 1},
		{1, 2, 3, 1, 2, 3},
	}

	k := []int{
		3,
		1,
		2,
	}

	ans := []bool{
		true,
		true,
		false,
	}

	for i := 0; i < len(nums); i++ {
		result := ContainsNearbyDuplicate(nums[i], k[i])
		if result != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
