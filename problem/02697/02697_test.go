package _02697

import "testing"

func TestSampe(t *testing.T) {
	s := []string{
		"egcfe",
		"abcd",
		"seven",
	}

	ans := []string{
		"efcfe",
		"abba",
		"neven",
	}

	for i, a := range ans {
		if result := makeSmallestPalindrome(s[i]); a != result {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
