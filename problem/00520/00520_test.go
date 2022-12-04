package _00520

import "testing"

func TestSample(t *testing.T) {
	words := []string{
		"USA",
		"FlaG",
		"usa",
		"Usa",
	}

	ans := []bool{
		true,
		false,
		true,
		true,
	}

	for i := 0; i < len(words); i++ {
		result := detectCapitalUse(words[i])
		if result != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
