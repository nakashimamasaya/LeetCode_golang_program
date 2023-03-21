package _02595

// 2595. Number of Even and Odd Bits
// https://leetcode.com/problems/number-of-even-and-odd-bits/

func evenOddBit(n int) []int {
	ans := []int{0, 0}
	for i := 0; n > 0; i, n = i+1, n>>1 {
		ans[i%2] += n % 2
	}
	return ans
}
