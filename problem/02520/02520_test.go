package _02520

import "testing"

func TestSample(t *testing.T) {
	nums := []int{
		7,
		121,
		1248,
	}

	ans := []int{
		1,
		2,
		4,
	}

	for i, n := range nums {
		if result := countDigits(n); result != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
