package _01071

import "testing"

func TestSample(t *testing.T) {
	str1 := []string{
		"ABCABC",
		"ABABAB",
		"LEET",
	}

	str2 := []string{
		"ABC",
		"ABAB",
		"CODE",
	}

	ans := []string{
		"ABC",
		"AB",
		"",
	}

	for i, a := range ans {
		if result := gcdOfStrings(str1[i], str2[i]); result != a {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
