package _01732

import "testing"

func TestSample(t *testing.T) {
	gain := [][]int{
		{-5, 1, 5, 0, -7},
		{-4, -3, -2, -1, 4, 3, 2},
	}

	ans := []int{
		1,
		0,
	}

	for i, a := range ans {
		if result := largestAltitude(gain[i]); a != result {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
