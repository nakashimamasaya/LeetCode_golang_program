package _00392

import "testing"

func TestSample(t *testing.T) {
	s := []string{
		"abc",
		"axc",
	}
	ts := []string{
		"ahbgdc",
		"ahbgdc",
	}

	ans := []bool{
		true,
		false,
	}

	for i, a := range ans {
		if result := isSubsequence(s[i], ts[i]); a != result {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}

}
