package _00029

import "testing"

func TestSample(t *testing.T) {
	dividend := []int{
		10,
		7,
	}

	divisor := []int{
		3,
		-3,
	}

	ans := []int{
		3,
		-2,
		-2147483647,
	}

	for i, d := range dividend {
		if result := divide(d, divisor[i]); result != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
