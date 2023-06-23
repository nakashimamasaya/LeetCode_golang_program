package _00334

import "testing"

func TestSample(t *testing.T) {
	nums := [][]int{
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{2, 1, 5, 0, 4, 6},
	}

	ans := []bool{
		true,
		false,
		true,
	}

	for i, a := range ans {
		if result := increasingTriplet(nums[i]); a != result {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}

}
