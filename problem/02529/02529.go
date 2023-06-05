package _02529

import "math"

// 2529. Maximum Count of Positive Integer and Negative Integer
// https://leetcode.com/problems/maximum-count-of-positive-integer-and-negative-integer/

func maximumCount(nums []int) int {
	ans := 0
	for i, n := range nums {
		if n > 0 {
			ans = int(math.Max(float64(ans), float64(len(nums)-i)))
			break
		}
		if n == 0 {
			continue
		}
		ans++
	}
	return ans
}
