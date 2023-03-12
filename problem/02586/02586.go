package _02586

// 2586. Count the Number of Vowel Strings in Range
// https://leetcode.com/problems/count-the-number-of-vowel-strings-in-range/

func vowelStrings(words []string, left int, right int) int {
	ans := 0
	for i := left; i <= right; i++ {
		if isVowel(words[i][0]) && isVowel(words[i][len(words[i])-1]) {
			ans++
		}
	}

	return ans
}

func isVowel(s byte) bool {
	const A_ASCII, I_ASCII, U_ASCII, E_ASCII, O_ASCII = 97, 105, 117, 101, 111
	return s == A_ASCII ||
		s == I_ASCII ||
		s == U_ASCII ||
		s == E_ASCII ||
		s == O_ASCII
}
