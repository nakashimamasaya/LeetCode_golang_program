package _0088

import (
	"testing"
)

func TestSample1(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)

	ans := []int{1, 2, 2, 3, 5, 6}
	for i := 0; i < len(nums1); i++ {
		if ans[i] != nums1[i] {
			t.Error("sample1 miss")
		}
	}
}

func TestSample2(t *testing.T) {
	nums1 := []int{1}
	merge(nums1, 1, []int{}, 0)

	if len(nums1) != 1 || nums1[0] != 1 {
		t.Error("sample2 miss")
	}
}

func TestSample3(t *testing.T) {
	nums1 := make([]int, 1)
	nums2 := []int{1}
	merge(nums1, 0, nums2, 1)

	if len(nums1) != 1 || nums1[0] != 1 {
		t.Error("sample3 miss")
	}
}
