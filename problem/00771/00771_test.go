package _00771

import "testing"

func TestSample(t *testing.T) {
	jewels := []string{
		"aA",
		"z",
	}

	stones := []string{
		"aAAbbbb",
		"ZZ",
	}

	ans := []int{
		3,
		0,
	}

	for i, a := range ans {
		if result := numJewelsInStones(jewels[i], stones[i]); result != a {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
