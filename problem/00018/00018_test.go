package _00018

import (
	"reflect"
	"sort"
	"testing"
)

func TestSample(t *testing.T) {
	nums := [][]int{
		{1, 0, -1, 0, -2, 2},
		{2, 2, 2, 2},
	}

	target := []int{
		0,
		8,
	}

	ans := [][][]int{
		{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		{{2, 2, 2, 2}},
	}

	for i, a := range ans {
		result := fourSum(nums[i], target[i])
		for _, v := range result {
			sort.Ints(v)
			flag := false
			for _, vv := range a {
				sort.Ints(vv)
				flag = flag || reflect.DeepEqual(v, vv)
			}
			if !flag {
				t.Errorf("miss sample %d\n return %v\n answer %v", i, result, a)
			}
		}
	}
}
