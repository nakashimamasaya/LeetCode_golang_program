package _02574

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	nums := [][]int{
		{10, 4, 8, 3},
		{1},
	}

	ans := [][]int{
		{15, 1, 11, 22},
		{0},
	}

	for i, a := range ans {
		if result := leftRigthDifference(nums[i]); !reflect.DeepEqual(a, result) {
			t.Errorf("miss sample %d\n return %v\n answer %v", i, result, a)
		}
	}
}
