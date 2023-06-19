package _00034

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	nums := [][]int{
		{5, 7, 7, 8, 8, 10},
		{5, 7, 7, 8, 8, 10},
		{},
	}

	target := []int{
		8,
		6,
		0,
	}

	ans := [][]int{
		{3, 4},
		{-1, -1},
		{-1, -1},
	}

	for i, a := range ans {
		if result := searchRange(nums[i], target[i]); !reflect.DeepEqual(result, a) {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
