package _00005

import "testing"

func TestSample(t *testing.T) {
	s := []string{
		"babad",
		"cbbd",
	}

	ans := []string{
		"bab",
		"bb",
	}

	for i, a := range ans {
		if result := longestPalindrome(s[i]); a != result {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
