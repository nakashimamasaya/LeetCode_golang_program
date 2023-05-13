package _01967

import "testing"

func TestSample(t *testing.T) {
	patterns := [][]string{
		{"a", "abc", "bc", "d"},
		{"a", "b", "c"},
		{"a", "a", "a"},
	}

	word := []string{
		"abc",
		"aaaaabbbbb",
		"ab",
	}

	ans := []int{
		3,
		2,
		3,
	}

	for i, a := range ans {
		if result := numOfStrings(patterns[i], word[i]); result != a {
			t.Errorf("miss sample %d\n return %v\n answer %v", i, result, a)
		}
	}
}
