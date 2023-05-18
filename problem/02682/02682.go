package _02682

// 2682. Find the Losers of the Circular Game
// https://leetcode.com/problems/find-the-losers-of-the-circular-game/

func circularGameLosers(n int, k int) []int {
	hand := make([]bool, n)

	for i, now := 1, 0; !hand[now%n]; i, now, hand[now%n] = i+1, now+i*k, true {
	}

	ans := []int{}
	for i, h := range hand {
		if !h {
			ans = append(ans, i+1)
		}
	}
	return ans
}
