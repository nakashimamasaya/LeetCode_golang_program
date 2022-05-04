package _01046

import "testing"

func TestSample(t *testing.T) {
	stones := [][]int{
		{2, 7, 4, 1, 8, 1},
		{1},
	}

	ans := []int{
		1,
		1,
	}
	for i := 0; i < len(stones); i++ {
		result := LastStoneWeight(stones[i])
		if result != ans[i] {
			t.Errorf("miss sample %d\n return %d\n answer %d", i+1, result, ans[i])
		}
	}
}
