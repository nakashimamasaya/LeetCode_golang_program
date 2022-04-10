package _0070

// 70. Climbing Stairs
// https://leetcode.com/problems/climbing-stairs/

func ClimbStairs(n int) int {
	ans := 1
	for i, fibo := 1, 1; i < n; i++ {
		tmp := fibo
		fibo, ans = ans, ans+tmp
	}
	return ans
}
