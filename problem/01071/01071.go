package _01071

import (
	"strings"
)

// 1071. Greatest Common Divisor of Strings
// https://leetcode.com/problems/greatest-common-divisor-of-strings/?envType=study-plan-v2&id=leetcode-75

func gcdOfStrings(str1 string, str2 string) string {
	ans := str1

	for ; len(ans) > 0 && !(checkZeroStrings(strings.Split(str1, ans)) && checkZeroStrings(strings.Split(str2, ans))); ans = ans[0 : len(ans)-1] {
	}
	return ans
}

func checkZeroStrings(str []string) bool {
	flag := true
	for i := 0; i < len(str) && flag; i++ {
		flag = len(str[i]) == 0
	}
	return flag
}
