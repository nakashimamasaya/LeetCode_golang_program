package _0017

// 17. Letter Combinations of a Phone Number
// https://leetcode.com/problems/letter-combinations-of-a-phone-number/

func LetterCombinations(digits string) []string {
	return letterCombinations(digits)
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	phone := map[string][]byte{
		"":  {0, 0},
		"2": {97, 3},
		"3": {100, 3},
		"4": {103, 3},
		"5": {106, 3},
		"6": {109, 3},
		"7": {112, 4},
		"8": {116, 3},
		"9": {119, 4},
	}
	ans := []string{}

	p := phone[string(digits[0])]
	for i := 0; i < int(p[1]); i++ {
		ans = append(ans, string(p[0]+byte(i)))
	}

	for i := 1; i < len(digits); i++ {
		p := phone[string(digits[i])]
		rl := []string{}
		for _, l := range ans {
			for j := 0; j < int(p[1]); j++ {
				rl = append(rl, l+string(p[0]+byte(j)))
			}
		}
		ans = rl
	}
	return ans
}
