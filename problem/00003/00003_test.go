package _00003

import "testing"

func TestSample(t *testing.T) {
	s := []string{
		"abcabcbb",
		"bbbbb",
		"pwwkew",
	}

	ans := []int{
		3,
		1,
		3,
	}

	for i, a := range ans {
		if result := lengthOfLongestSubstring(s[i]); result != a {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
