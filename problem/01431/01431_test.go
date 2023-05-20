package _01431

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	candies := [][]int{
		{2, 3, 5, 1, 3},
		{4, 2, 1, 1, 2},
		{12, 1, 12},
	}

	extraCandies := []int{
		3,
		1,
		10,
	}

	ans := [][]bool{
		{true, true, true, false, true},
		{true, false, false, false, false},
		{true, false, true},
	}

	for i, a := range ans {
		if result := kidsWithCandies(candies[i], extraCandies[i]); !reflect.DeepEqual(a, result) {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}

}
