package _0053

import (
	"math"
)

// 53. Maximum Subarray
// https://leetcode.com/problems/maximum-subarray/

func MaxSubArray(nums []int) int {
	ans := float64(nums[0])

	for i, memo := 0, 0.0; i < len(nums); i++ {
		memo = math.Max(float64(nums[i]), memo+float64(nums[i]))
		ans = math.Max(ans, memo)
	}
	return int(ans)
}
