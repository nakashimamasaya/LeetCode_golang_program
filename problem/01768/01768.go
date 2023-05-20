package _01768

// 1768. Merge Strings Alternately
// https://leetcode.com/problems/merge-strings-alternately/?envType=study-plan-v2&id=leetcode-75

func mergeAlternately(word1 string, word2 string) string {
	ans := []byte{}

	minLen, lemainWord := 0, ""

	if len(word1) <= len(word2) {
		minLen, lemainWord = len(word1), word2[len(word1):]
	} else {
		minLen, lemainWord = len(word2), word1[len(word2):]
	}

	for i := 0; i < minLen; i++ {
		ans = append(ans, word1[i], word2[i])
	}

	return string(ans) + lemainWord
}
