package _02595

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	n := []int{
		17,
		2,
	}

	ans := [][]int{
		{2, 0},
		{0, 1},
	}

	for i := 0; i < len(n); i++ {
		if result := evenOddBit(n[i]); !reflect.DeepEqual(ans[i], result) {
			t.Errorf("miss sample %d\n return %v\n answer %v", i, result, ans[i])

		}
	}
}
