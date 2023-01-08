package _00029

// 29. Divide Two Integers
// https://leetcode.com/problems/divide-two-integers/

func divide(dividend int, divisor int) int {
	negative, ans := 1, 0
	if dividend < 0 {
		negative, dividend = negative*-1, dividend*-1
	}
	if divisor < 0 {
		negative, divisor = negative*-1, divisor*-1
	}

	for ; dividend >= divisor; ans++ {
		dividend = dividend - divisor
	}

	return ans * negative
}
