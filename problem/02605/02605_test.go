package _02605

import "testing"

func TestSample(t *testing.T) {
	nums1 := [][]int{
		{4, 1, 3},
		{3, 5, 2, 6},
	}

	nums2 := [][]int{
		{5, 7},
		{3, 1, 7},
	}

	ans := []int{
		15,
		3,
	}

	for i, a := range ans {
		if result := minNumber(nums1[i], nums2[i]); a != result {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
