package _00345

// 345. Reverse Vowels of a String
// https://leetcode.com/problems/reverse-vowels-of-a-string

func reverseVowels(s string) string {
	vowels := []string{}
	for i := 0; i < len(s); i++ {
		if isVowel(string(s[i])) {
			vowels = append(vowels, string(s[i]))
		}
	}

	ans, vowelLen := "", len(vowels)

	for i := 0; i < len(s) && vowelLen > 0; i++ {
		if isVowel(string(s[i])) {
			vowelLen--
			ans += string(vowels[vowelLen])
		} else {
			ans += string(s[i])
		}
	}

	return ans + s[len(ans):]
}

func isVowel(s string) bool {
	vowels := []string{
		"a", "i", "u", "e", "o",
		"A", "I", "U", "E", "O",
	}

	for _, v := range vowels {
		if v == s {
			return true
		}
	}

	return false
}
