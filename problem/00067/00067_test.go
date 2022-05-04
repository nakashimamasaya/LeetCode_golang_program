package _0067

import "testing"

func TestSample(t *testing.T) {
	a := []string{
		"11",
		"1010",
	}

	b := []string{
		"1",
		"1011",
	}

	ans := []string{
		"100",
		"10101",
	}

	for i := 0; i < len(a); i++ {
		if AddBinary(a[i], b[i]) != ans[i] {
			t.Errorf("miss sample %d", i+1)
		}
	}
}
