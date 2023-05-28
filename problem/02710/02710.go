package _02710

// 2710. Remove Trailing Zeros From a String
// https://leetcode.com/problems/remove-trailing-zeros-from-a-string/

func removeTrailingZeros(num string) string {
	for i := len(num) - 1; i > 0 && string(num[i]) == "0"; i-- {
		num = num[:i]
	}
	return num
}
