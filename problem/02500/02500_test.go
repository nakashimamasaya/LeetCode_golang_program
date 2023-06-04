package _02500

import "testing"

func TestSample(t *testing.T) {
	grid := [][][]int{
		{{1, 2, 4}, {3, 3, 1}},
		{{10}},
	}

	ans := []int{
		8,
		10,
	}

	for i, a := range ans {
		if result := deleteGreatestValue(grid[i]); a != result {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
