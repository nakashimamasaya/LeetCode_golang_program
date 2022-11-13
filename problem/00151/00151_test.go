package _00151

import "testing"

func TestSample(t *testing.T) {
	s := []string{
		"the sky is blue",
		"  hello world  ",
		"a good   example",
	}

	ans := []string{
		"blue is sky the",
		"world hello",
		"example good a",
	}

	for i := 0; i < len(s); i++ {
		result := reverseWords(s[i])
		if result != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
