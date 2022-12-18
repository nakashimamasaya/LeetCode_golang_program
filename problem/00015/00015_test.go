package _00015

import "testing"

func TestSample(t *testing.T) {
	nums := [][]int{
		{-1, 0, 1, 2, -1, -4},
		{0, 1, 1},
		{0, 0, 0},
	}

	ans := [][][]int{
		{{-1, -1, 2}, {-1, 0, 1}},
		{},
		{{0, 0, 0}},
	}

	for i := 0; i < len(nums); i++ {
		result := threeSum(nums[i])
		if !contains(result, ans[i]) {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}

func contains(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
