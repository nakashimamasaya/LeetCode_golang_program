package _00075

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	nums := [][]int{
		{2, 0, 2, 1, 1, 0},
		{2, 0, 1},
	}

	ans := [][]int{
		{0, 0, 1, 1, 2, 2},
		{0, 1, 2},
	}

	for i, n := range nums {
		if sortColors(n); !reflect.DeepEqual(n, ans[i]) {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, n, ans[i])
		}
	}
}
