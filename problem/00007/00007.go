package _00007

import (
	"math"
)

// 7. Reverse Integer
// https://leetcode.com/problems/reverse-integer/

func reverse(x int) int {
	ans, negative := 0, 1
	if x < 0 {
		negative, x = -1, x*-1
	}

	for n := x; n > 0; n /= 10 {
		ans = ans*10 + n%10
	}
	ans *= negative

	if ans < math.MinInt32 || ans > math.MaxInt32 {
		return 0
	}
	return ans
}
