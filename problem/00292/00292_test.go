package _00292

import "testing"

func TestSample(t *testing.T) {
	n := []int{
		4,
		1,
		2,
	}

	ans := []bool{
		false,
		true,
		true,
	}

	for i, a := range ans {
		if result := canWinNim(n[i]); result != a {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
