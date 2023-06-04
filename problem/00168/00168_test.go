package _00168

import "testing"

func TestSample(t *testing.T) {
	columnNumber := []int{
		1,
		28,
		701,
	}

	ans := []string{
		"A",
		"AB",
		"ZY",
	}

	for i, a := range ans {
		if result := convertToTitle(columnNumber[i]); a != result {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
