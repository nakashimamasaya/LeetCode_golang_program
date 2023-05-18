package _02678

import "testing"

func TestSample(t *testing.T) {
	details := [][]string{
		{"7868190130M7522", "5303914400F9211", "9273338290F4010"},
		{"1313579440F2036", "2921522980M5644"},
	}

	ans := []int{
		2,
		0,
	}

	for i, a := range ans {
		if result := countSeniors(details[i]); result != a {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
