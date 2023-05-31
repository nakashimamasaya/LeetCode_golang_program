package _02515

import "testing"

func TestSample(t *testing.T) {
	words := [][]string{
		{"hello", "i", "am", "leetcode", "hello"},
		{"a", "b", "leetcode"},
		{"i", "eat", "leetcode"},
	}

	target := []string{
		"hello",
		"leetcode",
		"ate",
	}

	startIndex := []int{
		1,
		0,
		0,
	}

	ans := []int{
		1,
		1,
		-1,
	}

	for i, a := range ans {
		if result := closetTarget(words[i], target[i], startIndex[i]); a != result {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
