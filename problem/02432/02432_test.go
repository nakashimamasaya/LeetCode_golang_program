package _02432

import "testing"

func TestSample(t *testing.T) {
	n := []int{
		10,
		26,
		2,
	}

	logs := [][][]int{
		{{0, 3}, {2, 5}, {0, 9}, {1, 15}},
		{{1, 1}, {3, 7}, {2, 12}, {7, 17}},
		{{0, 10}, {1, 20}},
	}

	ans := []int{
		1,
		3,
		0,
	}

	for i, a := range ans {
		if result := hardestWorker(n[i], logs[i]); result != a {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
