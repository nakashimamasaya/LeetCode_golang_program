package _00016

import "testing"

func TestSample(t *testing.T) {
	nums := [][]int{
		{-1, 2, 1, -4},
		{0, 0, 0},
	}

	target := []int{
		1,
		1,
	}

	ans := []int{
		2,
		0,
	}

	for i, num := range nums {
		result := threeSumClosest(num, target[i])
		if result != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
