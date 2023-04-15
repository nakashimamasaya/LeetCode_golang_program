package _02404

import "testing"

func TestSample(t *testing.T) {
	nums := [][]int{
		{0, 1, 2, 2, 4, 4, 1},
		{4, 4, 4, 9, 2, 4},
		{29, 47, 21, 41, 13, 37, 25, 7},
	}

	ans := []int{
		2,
		4,
		-1,
	}

	for i, a := range ans {
		if result := mostFrequentEven(nums[i]); result != a {
			t.Errorf("miss sample %d\n return %v\n answer %v", i, result, a)
		}
	}
}
