package _0069

import "testing"

func TestSample(t *testing.T) {
	x := []int{
		4,
		8,
	}

	ans := []int{
		2,
		2,
	}

	for i := 0; i < len(x); i++ {
		result := MySqrt(x[i])
		if result != ans[i] {
			t.Errorf("miss sample %d\n return %d\n ans %d", i+1, result, ans[i])
		}
	}
}
