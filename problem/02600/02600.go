package _02600

import "math"

// 2600. K Items With the Maximum Sum
// https://leetcode.com/problems/k-items-with-the-maximum-sum/

func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	return int(math.Min(float64(numOnes), float64(k)) - math.Max(0, float64(k-numOnes-numZeros)))
}
