package _02605

import (
	"math"
	"sort"
)

// 2605. Form Smallest Number From Two Digit Arrays
// https://leetcode.com/problems/form-smallest-number-from-two-digit-arrays/

func minNumber(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	ans := 0
	if nums1[0] < nums2[0] {
		ans = nums1[0]*10 + nums2[0]
	} else {
		ans = nums2[0]*10 + nums1[0]
	}

	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			if num1 == num2 {
				ans = int(math.Min(
					float64(ans),
					float64(num1),
				))
				break
			}
		}
	}

	return ans
}
