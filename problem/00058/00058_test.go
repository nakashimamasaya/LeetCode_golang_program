package _0058

import "testing"

func TestSample(t *testing.T) {
	s := []string{
		"Hello World",
		"   fly me   to   the moon  ",
		"luffy is still joyboy",
	}
	ans := []int{
		5,
		4,
		6,
	}

	for i := 0; i < len(s); i++ {
		if LengthOfLastWord(s[i]) != ans[i] {
			t.Errorf("miss sammple %d", i+1)
		}
	}
}
