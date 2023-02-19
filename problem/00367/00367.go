package _00367

import "math"

// 367. Valid Perfect Square
// https://leetcode.com/problems/valid-perfect-square/

func isPerfectSquare(num int) bool {
	n := int(math.Sqrt(float64(num)))

	return math.Sqrt(float64(num)) == float64(n)
}
