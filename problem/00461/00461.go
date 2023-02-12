package _00461

// 461. Hamming Distance
// https://leetcode.com/problems/hamming-distance/

func hammingDistance(x int, y int) int {
	ans := 0
	for x ^= y; x > 0; ans, x = ans+1, x&(x-1) {
	}
	return ans
}
