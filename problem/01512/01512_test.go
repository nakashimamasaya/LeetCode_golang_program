package _01512

import "testing"

func TestSample(t *testing.T) {
	nums := [][]int{
		{1, 2, 3, 1, 1, 3},
		{1, 1, 1, 1},
		{1, 2, 3},
	}

	ans := []int{
		4,
		6,
		0,
	}

	for i, a := range ans {
		if result := numIdenticalPairs(nums[i]); result != a {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
