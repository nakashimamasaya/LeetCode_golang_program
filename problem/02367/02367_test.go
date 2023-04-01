package _02367

import "testing"

func TestSample(t *testing.T) {
	nums := [][]int{
		{0, 1, 4, 6, 7, 10},
		{4, 5, 6, 7, 8, 9},
	}

	diff := []int{
		3,
		2,
	}

	ans := []int{
		2,
		2,
	}

	for i := 0; i < len(nums); i++ {
		if result := arithmeticTriplets(nums[i], diff[i]); ans[i] != result {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
