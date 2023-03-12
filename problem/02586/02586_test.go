package _02586

import "testing"

func TestSample(t *testing.T) {
	words := [][]string{
		{"are", "amy", "u"},
		{"hey", "aeo", "mu", "ooo", "artro"},
	}
	left := []int{
		0,
		1,
	}
	right := []int{
		2,
		4,
	}

	ans := []int{
		2,
		3,
	}

	for i := 0; i < len(words); i++ {
		if result := vowelStrings(words[i], left[i], right[i]); ans[i] != result {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
