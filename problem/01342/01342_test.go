package _01342

import "testing"

func TestSample(t *testing.T) {
	num := []int{
		14,
		8,
		123,
	}

	ans := []int{
		6,
		4,
		12,
	}

	for i := 0; i < len(ans); i++ {
		if result := numberOfSteps(num[i]); ans[i] != result {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
