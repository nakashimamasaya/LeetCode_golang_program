package _02582

import "testing"

func TestSample(t *testing.T) {
	n := []int{
		4,
		3,
	}

	time := []int{
		5,
		2,
	}

	ans := []int{
		2,
		3,
	}

	for i := 0; i < len(ans); i++ {
		if result := passThePillow(n[i], time[i]); ans[i] != result {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
