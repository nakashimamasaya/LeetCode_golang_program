package _00263

// 263. Ugly Number
// https://leetcode.com/problems/ugly-number/

func isUgly(n int) bool {
	primeNum := []int{2, 3, 5}
	for _, pn := range primeNum {
		for ; n%pn == 0 && n > 0; n /= pn {
		}
	}
	return n == 1
}
