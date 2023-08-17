package _02810

import "testing"

func TestSample(t *testing.T) {
	s := []string{
		"string",
		"poiinter",
	}

	ans := []string{
		"rtsng",
		"ponter",
	}

	for i, a := range ans {
		if result := finalString(s[i]); a != result {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
