package _02529

import "testing"

func TestSample(t *testing.T) {
	nums := [][]int{
		{-2, -1, -1, 1, 2, 3},
		{-3, -2, -1, 0, 0, 1, 2},
		{5, 20, 66, 1314},
	}

	ans := []int{
		3,
		3,
		4,
	}

	for i, a := range ans {
		if result := maximumCount(nums[i]); result != a {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
