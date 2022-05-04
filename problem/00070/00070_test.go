package _0070

import "testing"

func TestSample(t *testing.T) {
	n := []int{
		2,
		3,
	}

	ans := []int{
		2,
		3,
	}

	for i := 0; i < len(n); i++ {
		result := ClimbStairs(n[i])
		if result != ans[i] {
			t.Errorf("miss sample %d\n return %d \n answer %d", i+1, result, ans[i])
		}
	}
}
