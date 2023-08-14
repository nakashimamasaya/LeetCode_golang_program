package _00004

import "testing"

func TestSample(t *testing.T) {
	nums1 := [][]int{
		{1, 3},
		{1, 2},
	}

	nums2 := [][]int{
		{2},
		{3, 4},
	}

	ans := []float64{
		2.00000,
		2.50000,
	}

	for i, a := range ans {
		if result := findMedianSortedArrays(nums1[i], nums2[i]); result != a {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
