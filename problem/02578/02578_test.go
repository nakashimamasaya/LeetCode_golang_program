package _02578

import "testing"

func TestSample(t *testing.T) {
	num := []int{
		4325,
		687,
	}

	ans := []int{
		59,
		75,
	}

	for i := 0; i < len(ans); i++ {
		if result := splitNum(num[i]); ans[i] != result {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
