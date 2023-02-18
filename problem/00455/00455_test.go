package _00455

import "testing"

func TestSample(t *testing.T) {
	g := [][]int{
		{1, 2, 3},
		{1, 2},
	}

	s := [][]int{
		{1, 1},
		{1, 2, 3},
	}

	ans := []int{
		1,
		2,
	}

	for i := 0; i < len(g); i++ {
		if result := findContentChildren(g[i], s[i]); result != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
