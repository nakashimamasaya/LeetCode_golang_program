package _00441

// 441. Arranging Coins
// https://leetcode.com/problems/arranging-coins/

func arrangeCoins(n int) int {

	ans := 0
	for i := 1; n >= i; i, ans = i+1, ans+1 {
		n -= i
	}
	return ans
}
