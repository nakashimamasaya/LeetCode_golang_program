package _02399

import "testing"

func TestSample(t *testing.T) {
	s := []string{
		"abaccb",
		"aa",
	}

	distance := [][]int{
		{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	ans := []bool{
		true,
		false,
	}

	for i, a := range ans {
		if result := checkDistances(s[i], distance[i]); result != a {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
