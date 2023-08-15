package _00125

import "unicode"

// 125. Valid Palindrome
// https://leetcode.com/problems/valid-palindrome/

func isPalindrome(s string) bool {
	ss, lss := make([]rune, len(s)), 0

	for _, si := range s {
		if unicode.IsUpper(si) {
			ss[lss] = 'a' - 'A'
		}
		if unicode.IsNumber(si) || unicode.IsLower(si) || unicode.IsUpper(si) {
			ss[lss] += si
			lss++
		}
	}

	for i := 0; i < lss/2; i++ {
		if ss[i] != ss[lss-1-i] {
			return false
		}
	}
	return true
}
