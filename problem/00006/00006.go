package _00006

import (
	"strings"
)

// 6. Zigzag Conversion
// https://leetcode.com/problems/zigzag-conversion/

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	tmpS := make([]string, numRows)
	numRows--
	for i, c := range s {
		if i/numRows%2 == 0 {
			tmpS[i%numRows] += string(c)
		} else {
			tmpS[numRows-i%numRows] += string(c)
		}
	}
	return strings.Join(tmpS, "")
}
