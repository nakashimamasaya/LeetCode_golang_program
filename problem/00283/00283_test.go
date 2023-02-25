package _00283

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	nums := [][]int{
		{0, 1, 0, 3, 12},
		{0},
	}

	ans := [][]int{
		{1, 3, 12, 0, 0},
		{0},
	}

	for i, n := range nums {
		moveZeroes(n)
		if !reflect.DeepEqual(n, ans[i]) {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, n, ans[i])
		}
	}
}
