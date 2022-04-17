package _0171

import (
	math "math"
)

// 171. Excel Sheet Column Number
// https://leetcode.com/problems/excel-sheet-column-number/

func TitleToNumber(columnTitle string) int {
	ans := 0
	for i := 0; i < len(columnTitle); i++ {
		ans += int(columnTitle[i]-64) * int(math.Pow(26, float64(len(columnTitle)-i-1)))
	}
	return ans
}
