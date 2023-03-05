package _00290

import "testing"

func TestSample(t *testing.T) {
	pattern := []string{
		"abba",
		"abba",
		"aaaa",
	}

	s := []string{
		"dog cat cat dog",
		"dog cat cat fish",
		"dog cat cat dog",
	}

	ans := []bool{
		true,
		false,
		false,
	}

	for i := 0; i < len(pattern); i++ {
		if result := wordPattern(pattern[i], s[i]); result != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
