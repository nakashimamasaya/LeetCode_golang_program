package _00033

import "testing"

func TestSample(t *testing.T) {
	nums := [][]int{
		{4, 5, 6, 7, 0, 1, 2},
		{4, 5, 6, 7, 0, 1, 2},
		{1},
	}

	target := []int{
		0,
		3,
		0,
	}

	ans := []int{
		4,
		-1,
		-1,
	}

	for i, a := range ans {
		if result := search(nums[i], target[i]); result != a {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
