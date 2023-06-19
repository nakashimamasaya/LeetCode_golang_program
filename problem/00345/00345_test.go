package _00345

import "testing"

func TestSample(t *testing.T) {
	s := []string{
		"hello",
		"leetcode",
	}

	ans := []string{
		"holle",
		"leotcede",
	}

	for i, a := range ans {
		if result := reverseVowels(s[i]); result != a {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
