package _0011

import "testing"

func TestSample(t *testing.T) {
	height := [][]int{
		{1, 8, 6, 2, 5, 4, 8, 3, 7},
		{1, 1},
	}

	ans := []int{
		49,
		1,
	}

	for i := 0; i < len(ans); i++ {
		result := maxArea(height[i])
		if result != ans[i] {
			t.Errorf("miss sample %v\n return %v \n answer %v", i+1, result, ans[i])
		}
	}
}
