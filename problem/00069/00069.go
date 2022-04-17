package _0069

// 69. Sqrt(x)
// https://leetcode.com/problems/sqrtx/

func MySqrt(x int) int {
	for i := 0; ; i++ {
		if i*i > x {
			return i - 1
		}
	}
}
