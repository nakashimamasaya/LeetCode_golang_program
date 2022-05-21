package _0017

import (
	"testing"
)

func TestSample(t *testing.T) {
	digits := []string{
		"23",
		"",
		"2",
	}

	ans := [][]string{
		{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		{},
		{"a", "b", "c"},
	}

	for i := 0; i < len(ans); i++ {
		result := letterCombinations(digits[i])
		for k, v := range ans[i] {
			if v != result[k] {
				t.Errorf("miss sample %d\n return %v \n answer %v", i+1, result, ans[i])
				break
			}
		}
	}

}
