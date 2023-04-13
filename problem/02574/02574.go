package _02574

import "math"

// 2574. Left and Right Sum Differences
// https://leetcode.com/problems/left-and-right-sum-differences/

func leftRigthDifference(nums []int) []int {
	ans := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i-1] + nums[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		ans[i] = int(math.Abs(float64(ans[i] - nums[i+1])))
		nums[i] += nums[i+1]
	}
	return ans
}
