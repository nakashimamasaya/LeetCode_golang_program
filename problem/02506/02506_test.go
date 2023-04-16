package _02506

import "testing"

func TestSample(t *testing.T) {
	words := [][]string{
		{"aba", "aabb", "abcd", "bac", "aabc"},
		{"aabb", "ab", "ba"},
		{"nba", "cba", "dba"},
	}

	ans := []int{
		2,
		3,
		0,
	}

	for i, a := range ans {
		if result := similarPairs(words[i]); result != a {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
