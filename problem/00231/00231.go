package _00231

// 231. Power of Two
// https://leetcode.com/problems/power-of-two/

func isPowerOfTwo(n int) bool {
	return n&(n-1) == 0 && n != 0
}
