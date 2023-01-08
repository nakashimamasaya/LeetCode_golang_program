package _00008

import (
	"math"
	"strings"
)

// 8. String to Integer (atoi)
// https://leetcode.com/problems/string-to-integer-atoi/

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	negativeFlag := s[0] == '-'

	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
	}

	ans := 0
	for i := 0; i < len(s) && s[i] >= 48 && s[i] <= 57 && math.MaxInt32 >= ans; i++ {
		ans = ans*10 + int(s[i]-48)
	}

	if negativeFlag {
		ans *= -1
	}
	if math.MinInt32 > ans {
		ans = math.MinInt32
	} else if math.MaxInt32 < ans {
		ans = math.MaxInt32
	}

	return ans
}
