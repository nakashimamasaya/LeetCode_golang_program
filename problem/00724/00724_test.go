package _00724

import "testing"

func TestSample(t *testing.T) {
	nums := [][]int{
		{1, 7, 3, 6, 5, 6},
		{1, 2, 3},
		{2, 1, -1},
	}

	ans := []int{
		3,
		-1,
		0,
	}

	for i, a := range ans {
		if result := pivotIndex(nums[i]); a != result {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
