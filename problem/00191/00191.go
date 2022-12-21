package _00191

// 191. Number of 1 BitS
// https://leetcode.com/problems/number-of-1-bits/

func hammingWeight(num uint32) int {
	ans := 0
	for ; num != 0; num, ans = num>>1, ans+int(num%2) {
	}
	return ans
}
