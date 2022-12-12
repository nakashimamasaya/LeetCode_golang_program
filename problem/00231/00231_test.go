package _00231

import "testing"

func TestSample(t *testing.T) {
	n := []int{
		1,
		16,
		3,
		0,
	}

	ans := []bool{
		true,
		true,
		false,
		false,
	}

	for i := 0; i < len(n); i++ {
		result := isPowerOfTwo(n[i])
		if result != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
