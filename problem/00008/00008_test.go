package _00008

import "testing"

func TestSample(t *testing.T) {
	s := []string{
		"42",
		"   -42",
		"4193 with words",
	}

	ans := []int{
		42,
		-42,
		4193,
	}

	for i, str := range s {
		if result := myAtoi(str); result != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
