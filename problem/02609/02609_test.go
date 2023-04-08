package _02609

import "testing"

func TestSample(t *testing.T) {
	s := []string{
		"01000111",
		"00111",
		"111",
	}

	ans := []int{
		6,
		4,
		0,
	}

	for i, a := range ans {
		if result := findTheLongestBalancedSubstring(s[i]); a != result {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
