package _00383

import "testing"

func TestSample(t *testing.T) {
	ransomNote := []string{
		"a",
		"aa",
		"aa",
	}

	magazine := []string{
		"b",
		"ab",
		"aab",
	}

	ans := []bool{
		false,
		false,
		true,
	}

	for i := 0; i < len(ransomNote); i++ {
		if result := canConstruct(ransomNote[i], magazine[i]); result != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
