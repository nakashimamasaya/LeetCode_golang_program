package _00238

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	nums := [][]int{
		{1, 2, 3, 4},
		{-1, 1, 0, -3, 3},
	}

	ans := [][]int{
		{24, 12, 8, 6},
		{0, 0, 9, 0, 0},
	}

	for i, a := range ans {
		if result := productExceptSelf(nums[i]); !reflect.DeepEqual(result, a) {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
