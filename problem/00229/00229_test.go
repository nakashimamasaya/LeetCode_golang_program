package _00229

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	nums := [][]int{
		{3, 2, 3},
		{1},
		{1, 2},
	}

	ans := [][]int{
		{3},
		{1},
		{1, 2},
	}

	for i, n := range nums {
		if result := majorityElement(n); !reflect.DeepEqual(result, ans[i]) {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
