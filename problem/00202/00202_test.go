package _00202

import "testing"

func TestSample(t *testing.T) {
	n := []int{
		19,
		2,
	}

	ans := []bool{
		true,
		false,
	}

	for i, tmp := range n {
		result := isHappy(tmp)

		if result != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
