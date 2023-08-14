package _00004

import "sort"

// 4. Median of Two Sorted Arrays
// https://leetcode.com/problems/median-of-two-sorted-arrays/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	if len(nums)%2 == 0 {
		return float64(nums[len(nums)/2]+nums[len(nums)/2-1]) / 2
	}
	return float64(nums[len(nums)/2])
}
