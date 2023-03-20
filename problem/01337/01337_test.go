package _01337

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	mat := [][][]int{
		{
			{1, 1, 0, 0, 0},
			{1, 1, 1, 1, 0},
			{1, 0, 0, 0, 0},
			{1, 1, 0, 0, 0},
			{1, 1, 1, 1, 1},
		},
		{
			{1, 0, 0, 0},
			{1, 1, 1, 1},
			{1, 0, 0, 0},
			{1, 0, 0, 0},
		},
	}

	k := []int{
		3,
		2,
	}

	ans := [][]int{
		{2, 0, 3},
		{0, 2},
	}
	for i := 0; i < len(mat); i++ {
		if result := kWeakestRows(mat[i], k[i]); !reflect.DeepEqual(result, ans[i]) {
			t.Errorf("miss sample %d\n return %v\n answer %v", i, result, ans[i])
		}
	}

}
