package _00050

// 50. Pow(x, n)
// https://leetcode.com/problems/powx-n/

func myPow(x float64, n int) float64 {
	ans := 1.0
	if n < 0 {
		n, x = n*-1, 1/x
	}
	switch {
	case x == -1 && n%2 == 1:
		ans = -1
	case x == 1 || x == -1:
		ans = 1
	case x == 2 || x == -2:
		ans = float64(int(ans) << n)
	default:
		for ; n != 0; n-- {
			ans *= x
		}
	}
	return ans
}
