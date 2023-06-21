package _00441

import "testing"

func TestSample(t *testing.T) {
	n := []int{
		5,
		8,
	}

	ans := []int{
		2,
		3,
	}

	for i, a := range ans {
		if result := arrangeCoins(n[i]); result != a {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
