package _02506

// 2506. Count Pairs Of Similar Strings
// https://leetcode.com/problems/count-pairs-of-similar-strings/

func similarPairs(words []string) int {
	ans := 0
	wordsMap := map[string]map[rune]int{}

	for _, w := range words {
		wordsMap[w] = wordMap(w)
	}

	for i, wo := range words {
		for _, ww := range words[i+1:] {
			if similarPair(wordsMap[wo], wordsMap[ww]) {
				ans++
			}
		}
	}

	return ans
}

func wordMap(s string) map[rune]int {
	m := map[rune]int{}
	for _, ss := range s {
		m[ss]++
	}
	return m
}

func similarPair(aMap, bMap map[rune]int) bool {
	if len(aMap) != len(bMap) {
		return false
	}
	for k := range aMap {
		if bMap[k] == 0 {
			return false
		}
	}
	return true
}
