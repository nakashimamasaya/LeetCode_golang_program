package _02315

import "testing"

func TestSample(t *testing.T) {
	ss := []string{
		"l|*e*et|c**o|*de|",
		"iamprogrammer",
		"yo|uar|e**|b|e***au|tifu|l",
	}

	ans := []int{
		2,
		0,
		5,
	}

	for i, s := range ss {
		if result := countAsterisks(s); ans[i] != result {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
