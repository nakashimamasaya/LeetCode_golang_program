package _00605

import "testing"

func TestSample(t *testing.T) {
	flowerbed := [][]int{
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
	}

	n := []int{
		1,
		2,
	}

	ans := []bool{
		true,
		false,
	}

	for i, a := range ans {
		if result := canPlaceFlowers(flowerbed[i], n[i]); result != a {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
