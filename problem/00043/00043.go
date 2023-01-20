package _00043

import (
	"strconv"
	"strings"
)

// 43. Multiply Strings
// https://leetcode.com/problems/multiply-strings/

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	ans := "0"
	for i := len(num1) - 1; i >= 0; i-- {
		t, carry := strings.Repeat("0", len(num1)-i-1), 0
		for j := len(num2) - 1; j >= 0; j-- {
			cal := (num2[j]-48)*(num1[i]-48) + byte(carry)
			carry = int(cal) / 10
			t += strconv.Itoa(int(cal % 10))
		}

		if carry != 0 {
			t += strconv.Itoa(carry)
			carry = 0
		}

		t = reverse(t)
		tmp := ""
		for j := 0; j < len(ans) || j < len(t); j++ {
			cal := 0
			if j > len(ans)-1 {
				cal = carry + int(t[len(t)-1-j]) - 48
			} else if j > len(t)-1 {
				cal = carry + int(ans[len(ans)-1-j]) - 48
			} else {
				cal = carry + int(ans[len(ans)-1-j]+t[len(t)-1-j]) - 96
			}
			carry = int(cal) / 10
			tmp += strconv.Itoa(int(cal % 10))
		}
		if carry == 1 {
			tmp += "1"
		}
		ans = reverse(tmp)
	}

	return ans
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
