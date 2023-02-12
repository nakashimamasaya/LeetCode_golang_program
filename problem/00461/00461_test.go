package _00461

import "testing"

func TestSample(t *testing.T) {
	x := []int{
		1,
		3,
	}

	y := []int{
		4,
		1,
	}

	ans := []int{
		2,
		1,
	}

	for i := 0; i < len(x); i++ {
		if result := hammingDistance(x[i], y[i]); ans[i] != result {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
