package _02682

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	n := []int{
		5,
		4,
	}

	k := []int{
		2,
		4,
	}

	ans := [][]int{
		{4, 5},
		{2, 3, 4},
	}

	for i, a := range ans {
		if result := circularGameLosers(n[i], k[i]); !reflect.DeepEqual(a, result) {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
