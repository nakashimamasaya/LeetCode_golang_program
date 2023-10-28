package _02908

import "testing"

func TestSample(t *testing.T) {
	nums := [][]int{
		{8, 6, 1, 5, 3},
		{5, 4, 8, 7, 10, 2},
		{6, 5, 4, 3, 4, 5},
	}

	ans := []int{
		9,
		13,
		-1,
	}

	for i, a := range ans {
		if result := minimumSum(nums[i]); result != a {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
