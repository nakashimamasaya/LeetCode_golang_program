package _00005

// 5. Longest Palindromic Substring
// https://leetcode.com/problems/longest-palindromic-substring/

func longestPalindrome(s string) string {
	rs := reverse(s)
	for i := len(s); i > 1; i-- {
		for j := 0; j+i <= len(s); j++ {
			if s[j:j+i/2] == rs[len(s)-i-j:len(s)-i+i/2-j] {
				return s[j : j+i]
			}
		}
	}
	return string(s[0])
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
