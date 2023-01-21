package _00120

import "testing"

func TestSample(t *testing.T) {
	triangle := [][][]int{
		{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}},
		{{-10}},
	}

	ans := []int{
		11,
		-10,
	}

	for i, tr := range triangle {
		if result := minimumTotal(tr); ans[i] != result {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
