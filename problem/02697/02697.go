package _02697

// 2697. Lexicographically Smallest Palindrome
// https://leetcode.com/problems/lexicographically-smallest-palindrome/

func makeSmallestPalindrome(s string) string {
	runeS := []rune(s)
	for i := 0; i < len(s)/2; i++ {
		if s[i] > s[len(s)-1-i] {
			runeS[i] = runeS[len(s)-1-i]
		} else {
			runeS[len(s)-1-i] = runeS[i]
		}
	}
	return string(runeS)
}
