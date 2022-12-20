package _00016

import (
	"math"
	"sort"
)

// 16. 3Sum Closest
// https://leetcode.com/problems/3sum-closest/

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans := nums[0] + nums[1] + nums[2]
	abs := math.Abs(float64(target - ans))
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]
				sumAbs := math.Abs(float64(target - sum))
				if abs != math.Min(abs, sumAbs) {
					abs, ans = sumAbs, sum
				}
			}
		}
	}
	return ans
}
