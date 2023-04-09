package _02614

import "testing"

func TestSample(t *testing.T) {
	nums := [][][]int{
		{{1, 2, 3}, {5, 6, 7}, {9, 10, 11}},
		{{1, 2, 3}, {5, 17, 7}, {9, 11, 10}},
	}

	ans := []int{
		11,
		17,
	}

	for i, a := range ans {
		if result := diagonalPrime(nums[i]); result != a {
			t.Errorf("miss sample %d\n return %v\n answer %v", i, result, a)
		}
	}
}
