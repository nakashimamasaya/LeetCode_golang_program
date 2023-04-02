package _02300

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	spelles := [][]int{
		{5, 1, 3},
		{3, 1, 2},
	}

	potions := [][]int{
		{1, 2, 3, 4, 5},
		{8, 5, 8},
	}

	success := []int64{
		7,
		16,
	}

	ans := [][]int{
		{4, 0, 3},
		{2, 0, 2},
	}

	for i := 0; i < len(ans); i++ {
		if result := successfulPairs(spelles[i], potions[i], success[i]); !reflect.DeepEqual(ans[i], result) {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
