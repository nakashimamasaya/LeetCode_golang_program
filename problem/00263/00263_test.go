package _00263

import "testing"

func TestSample(t *testing.T) {
	ns := []int{
		6,
		1,
		14,
	}

	ans := []bool{
		true,
		true,
		false,
	}

	for i, n := range ns {
		if result := isUgly(n); result != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
