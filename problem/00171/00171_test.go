package _0171

import "testing"

func TestSample(t *testing.T) {
	columnTitle := []string{
		"A",
		"AB",
		"ZY",
	}

	ans := []int{
		1,
		28,
		701,
	}

	for i := 0; i < len(columnTitle); i++ {
		result := TitleToNumber(columnTitle[i])
		if result != ans[i] {
			t.Errorf("miss sample %d\nreturn %v\nanswer %v", i+1, result, ans[i])
		}
	}
}
