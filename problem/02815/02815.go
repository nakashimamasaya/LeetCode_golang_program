package _02815

import "math"

// 2815. Max Pair Sum in an Array
// https://leetcode.com/problems/max-pair-sum-in-an-array/

func maxSum(nums []int) int {
	digit, ans := [10]int{}, -1
	for _, n := range nums {
		d := maxDigit(n)
		if digit[d] > 0 {
			ans = int(math.Max(float64(ans), float64(n+digit[d])))
		}
		digit[d] = int(math.Max(float64(digit[d]), float64(n)))
	}

	return ans
}

func maxDigit(n int) int {
	result := 0
	for ; n > 0; n /= 10 {
		if n%10 > result {
			result = n % 10
		}
	}
	return result
}
