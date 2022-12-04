package _00520

import "unicode"

// 520. Detect Capital
// https://leetcode.com/problems/detect-capital/

func detectCapitalUse(word string) bool {
	r := []rune(word)
	if unicode.IsLower(r[0]) {
		return isAllLowerCase(r)
	}
	return isAllUpperCase(r) || isAllLowerCase(r[1:])
}

func isAllLowerCase(word []rune) bool {
	for _, s := range word {
		if !unicode.IsLower(s) {
			return false
		}
	}
	return true
}

func isAllUpperCase(word []rune) bool {
	for _, s := range word {
		if !unicode.IsUpper(s) {
			return false
		}
	}
	return true
}
