package _02710

import "testing"

func TestSample(t *testing.T) {
	num := []string{
		"51230100",
		"123",
	}

	ans := []string{
		"512301",
		"123",
	}

	for i, a := range ans {
		if result := removeTrailingZeros(num[i]); a != result {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
