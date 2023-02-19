package _00367

import "testing"

func TestSample(t *testing.T) {
	num := []int{
		16,
		14,
	}

	ans := []bool{
		true,
		false,
	}

	for i, n := range num {
		if result := isPerfectSquare(n); ans[i] != result {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
